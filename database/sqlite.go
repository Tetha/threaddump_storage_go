package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/tetha/threaddumpstorage-go/input"
	"github.com/tetha/threaddumpstorage-go/model"
)

/*SQLiteStore is an implementation of dataastor based on sqlite*/
type SQLiteStore struct {
	db *sql.DB
}

/*NewSQLiteStore opens a new connection to the database*/
func NewSQLiteStore(filename string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		return nil, errors.Wrap(err, "connection failed")
	}

	return &SQLiteStore{db}, nil
}

/*ListPagedThreadHeaders is used to paginate thread headers in a threaddump*/
func (store *SQLiteStore) ListPagedThreadHeaders(threaddumpID string, limit int, offset int) ([]model.JavaThreadHeader, error) {
	rows, err := store.db.Query(`SELECT id, name, java_id, is_daemon, prio, os_prio, tid, nid, native_thread_state, condition_address, java_thread_state, java_state_clarification
						         FROM java_threads
						         WHERE threaddump_id = ?
						         ORDER BY id
						         LIMIT ? OFFSET ?`, threaddumpID, limit, offset)
	if err != nil {
		return nil, err
	}

	threads := []model.JavaThreadHeader{}
	for rows.Next() {
		thread, err := decodeJavaThreadHeader(rows)
		if err != nil {
			return nil, err
		}
		threads = append(threads, *thread)
	}
	return threads, nil
}

/*ListAllThreadHeadersInDump implementation for sqlite*/
func (store *SQLiteStore) ListAllThreadHeadersInDump(threaddumpID string) ([]model.JavaThreadHeader, error) {
	rows, err := store.db.Query(`SELECT id, name, java_id, is_daemon, prio, os_prio, tid, nid, native_thread_state, condition_address, java_thread_state, java_state_clarification
						   FROM java_threads
						   WHERE threaddump_id = ?
						   ORDER BY id`, threaddumpID)
	if err != nil {
		log.Printf("Error with db query: %s", err)
		return nil, errors.New("Query error")
	}

	threads := []model.JavaThreadHeader{}
	for rows.Next() {
		thread, err := decodeJavaThreadHeader(rows)
		if err != nil {
			return nil, err
		}
		threads = append(threads, *thread)
	}
	return threads, nil
}

func decodeJavaThreadHeader(rows *sql.Rows) (*model.JavaThreadHeader, error) {
	var thread model.JavaThreadHeader

	var name sql.NullString
	var javaID sql.NullString
	// isDaemon is boolean, not nullable
	var prio sql.NullInt64
	var osPrio sql.NullInt64
	var tid sql.NullString
	var nid sql.NullString
	var nativeThreadState sql.NullString
	var conditionAddress sql.NullString
	var javaThreadState sql.NullString
	var javaThreadStateClarification sql.NullString

	err := rows.Scan(&thread.ID, &name, &javaID, &thread.IsDaemon, &prio, &osPrio, &tid, &nid, &nativeThreadState, &conditionAddress, &javaThreadState, &javaThreadStateClarification)
	if err != nil {
		log.Printf("Error with db query: %s", err)
		return nil, errors.Wrap(err, "Scan error")
	}

	if name.Valid {
		thread.Name = name.String
	}
	if javaID.Valid {
		thread.JavaID = javaID.String
	}
	if prio.Valid {
		thread.Prio = int(prio.Int64)
	} else {
		thread.Prio = -1
	}
	if osPrio.Valid {
		thread.OsPrio = int(osPrio.Int64)
	} else {
		thread.OsPrio = -1
	}
	if tid.Valid {
		thread.Tid = tid.String
	}
	if nid.Valid {
		thread.Nid = nid.String
	}
	if nativeThreadState.Valid {
		thread.NativeThreadState = nativeThreadState.String
	}
	if conditionAddress.Valid {
		thread.ConditionAddress = conditionAddress.String
	}
	if javaThreadState.Valid {
		thread.JavaThreadState = javaThreadState.String
	}
	if javaThreadStateClarification.Valid {
		thread.JavaStateClarification = javaThreadStateClarification.String
	}

	return &thread, nil
}

func (store *SQLiteStore) ListAllThreaddumps() ([]model.Threaddump, error) {
	rows, err := store.db.Query("SELECT id, application, host, upload_time FROM threaddumps")
	if err != nil {
		log.Printf("Error with db query: %s", err)
		return nil, err
	}

	dumps := []model.Threaddump{}
	for rows.Next() {
		var dump model.Threaddump
		err := rows.Scan(&dump.ID, &dump.Application, &dump.Host, &dump.Uploaded)
		if err != nil {
			log.Printf("Error scanning result rows: %s", err)
			return nil, err
		}
		dumps = append(dumps, dump)
	}
	return dumps, nil
}

func (store *SQLiteStore) StoreDump(application, host string, dump input.Threaddump) (string, error) {
	tx, err := store.db.Begin()
	if err != nil {
		return "", err
	}

	res, err := tx.Exec("INSERT INTO threaddumps (application, host, upload_time) VALUES (?, ?, ?)", application, host, time.Now())
	if err != nil {
		tx.Rollback()
		return "", errors.Wrap(err, "Unable to store threaddump header")
	}
	dumpID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return "", errors.Wrap(err, "Unable to get the ID of the dump")
	}
	for _, thread := range dump.Threads {
		res, err = tx.Exec("INSERT INTO java_threads (name, java_id, is_daemon, prio, os_prio, tid, nid, native_thread_state, condition_address, java_thread_state, java_state_clarification, threaddump_id)"+
			"VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			thread.Name, thread.ID, thread.IsDaemon, thread.Prio, thread.OsPrio, thread.Tid, thread.Nid, thread.ThreadState, thread.ConditionAddress, thread.JavaState, thread.JavaStateDetail,
			dumpID)
		if err != nil {
			tx.Rollback()
			return "", errors.Wrapf(err, "Error storing thread %v", thread)
		}

		threadIDInDB, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "Error getting thread id")
		}

		for idx, line := range thread.Stacktrace {
			baseQuery := "INSERT INTO stacktrace_lines (kind, line_number, lock_address, locked_class, lock_class, java_class, java_method, source_line, source_file, java_thread_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			switch line.Type {
			case input.Uninitialized:
				return "", fmt.Errorf("Uninitialized line found: %v", line)
			case input.WaitingLine:
				_, err = tx.Exec(baseQuery, line.Type, idx, line.LockAddress, "", line.Class, "", "", "", "", threadIDInDB)
			case input.BlockedLine:
				_, err = tx.Exec(baseQuery, line.Type, idx, line.LockAddress, line.Class, "", "", "", "", "", threadIDInDB)
			case input.LockedLine:
				_, err = tx.Exec(baseQuery, line.Type, idx, line.LockAddress, line.Class, "", "", "", "", "", threadIDInDB)
			case input.PositionLine:
				_, err = tx.Exec(baseQuery, line.Type, idx, "", "", "", line.Class, line.Method, line.SourceLine, line.SourceFile, threadIDInDB)
			case input.ParkedLine:
				_, err = tx.Exec(baseQuery, line.Type, idx, line.LockAddress, line.Class, "", "", "", "", "", threadIDInDB)
			}

			if err != nil {
				tx.Rollback()
				return "", errors.Wrapf(err, "Unable to store source line %v", line)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return "", errors.Wrap(err, "Unable to commit transaction")
	}
	return string(dumpID), nil
}
