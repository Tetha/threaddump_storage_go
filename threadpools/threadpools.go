package threadpools

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"html/template"
)

type javaThread struct {
	ID                     int
	Name                   string
	JavaID                 string
	IsDaemon               bool
	Prio                   int
	OsPrio                 int
	Tid                    string
	Nid                    string
	NativeThreadState      string
	ConditionAddress       string
	JavaThreadState        string
	JavaStateClarification string
}

type threadPoolDetection struct {
	ThreadPools    map[string][]javaThread
	UnknownThreads []javaThread
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

// ListThreadpools is an http handler which displays all detected threadpools
func ListThreadpools(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	threaddumpID := r.URL.Path[len("/threadpools/"):]
	if threaddumpID == "" {
		http.Error(w, "400: Missing threaddump id", 400)
		return
	}

	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Printf("Error opening database: %s", err)
		http.Error(w, "500: cannot open database", 500)
		return
	}
	threads, err := getThreadsFromDB(db, threaddumpID)
	if err != nil {
		http.Error(w, "500: cannot get threads from db", 500)
		return
	}

	detectedPools := figureOutThreadpools(threads)

	err = templates.ExecuteTemplate(w, "threadpool_list.html", detectedPools)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
	}

}

func figureOutThreadpools(threads []javaThread) (pools threadPoolDetection) {
	pools.ThreadPools = make(map[string][]javaThread)
	pools.UnknownThreads = threads

	findElasticsearchThreadpools(&pools)
	findNumberedThreads(&pools)
	return
}

func findNumberedThreads(pools *threadPoolDetection) {
	var newUnknownThreads []javaThread
	nameRegex := regexp.MustCompile(`(.+?)[-#]?\d+$`)
	for _, thread := range pools.UnknownThreads {
		if match := nameRegex.FindStringSubmatch(thread.Name); match != nil {
			readableName := match[1]
			pools.ThreadPools[readableName] = append(pools.ThreadPools[readableName], thread)
		} else {
			newUnknownThreads = append(newUnknownThreads, thread)
		}
	}
	pools.UnknownThreads = newUnknownThreads
}

func findElasticsearchThreadpools(pools *threadPoolDetection) {
	var newUnknownThreads []javaThread
	nameRegex := regexp.MustCompile(`^elasticsearch\[([^]]+)\]\[([^]]+)\]`)
	for _, thread := range pools.UnknownThreads {
		if match := nameRegex.FindStringSubmatch(thread.Name); match != nil {
			readableName := fmt.Sprintf("elasticsearch (instance=%s, pool=%s)", match[1], match[2])
			pools.ThreadPools[readableName] = append(pools.ThreadPools[readableName], thread)
		} else {
			newUnknownThreads = append(newUnknownThreads, thread)
		}
	}
	pools.UnknownThreads = newUnknownThreads
}

func getThreadsFromDB(db *sql.DB, threaddumpID string) ([]javaThread, error) {
	rows, err := db.Query(`SELECT id, name, java_id, is_daemon, prio, os_prio, tid, nid, native_thread_state, condition_address, java_thread_state, java_state_clarification
						   FROM java_threads
						   WHERE threaddump_id = ?
						   ORDER BY id`, threaddumpID)
	if err != nil {
		log.Printf("Error with db query: %s", err)
		return nil, errors.New("Query error")
	}

	threads := []javaThread{}
	for rows.Next() {
		var thread javaThread

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
