package threads

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

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

type pagedThreadList struct {
	ThreaddumpID string
	From         int
	Limit        int

	Threads []javaThread
}

func (list pagedThreadList) PreviousPageStart() int {
	return list.From - list.Limit
}

func (list pagedThreadList) NextPageStart() int {
	return list.From + list.Limit
}

// ListThreads displays a paged list of the threads in a threaddump
func ListThreads(w http.ResponseWriter, r *http.Request) {
	// TODO: validation

	r.ParseForm()
	threaddumpID := r.URL.Path[len("/threads/"):]
	if threaddumpID == "" {
		http.Error(w, "400: Missing threaddump id", 400)
		return
	}

	fromRaw := r.FormValue("from")
	if fromRaw == "" {
		fromRaw = "0"
	}
	from, err := strconv.Atoi(fromRaw)
	if err != nil {
		http.Error(w, "400: From is not valid int", 400)
		return
	}
	if from < 0 {
		http.Error(w, "400: from must be positive", 400)
		return
	}

	limitRaw := r.FormValue("limit")
	if limitRaw == "" {
		limitRaw = "20"
	}
	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		http.Error(w, "400: To is not valid int", 400)
		return
	}
	log.Printf("Listing from %d to %d + %d", from, from, limit)

	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name, java_id, is_daemon, prio, os_prio, tid, nid, native_thread_state, condition_address, java_thread_state, java_state_clarification
						   FROM java_threads
						   WHERE threaddump_id = ?
						   ORDER BY id
						   LIMIT ? OFFSET ?`, threaddumpID, limit, from)
	if err != nil {
		log.Printf("Error with db query: %s", err)
		http.Error(w, "500: Database Error", 500)
		return
	}

	threads := []javaThread{}
	for rows.Next() {
		var thread javaThread
		err := rows.Scan(&thread.ID, &thread.Name, &thread.JavaID, &thread.IsDaemon, &thread.Prio, &thread.OsPrio, &thread.Tid, &thread.Nid, &thread.NativeThreadState, &thread.ConditionAddress, &thread.JavaThreadState, &thread.JavaStateClarification)
		if err != nil {
			log.Printf("Error with db query: %s", err)
			http.Error(w, "500: Database Error", 500)
			return
		}
		threads = append(threads, thread)
	}
	templateInput := pagedThreadList{threaddumpID, from, limit, threads}
	err = templates.ExecuteTemplate(w, "threads_list.html", templateInput)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
	}
}
