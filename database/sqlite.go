package database

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
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

/*ListThreadHeadersInDump implementation for sqlite*/
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
			return nil, errors.New("Scan error")
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

		threads = append(threads, thread)
	}
	return threads, nil
}
