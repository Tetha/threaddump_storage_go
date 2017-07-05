package main

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"encoding/json"
	"net/http"
	"goji.io"
	"goji.io/pat"
	"github.com/rs/cors"
)

type Threaddump struct {
	Id int `json:"id"`
	Application string `json:"application"`
	Host string `json:"host"`
	Uploaded time.Time `json:"uploaded"`
}

type JavaThread struct {
	Id int `json:"id"`
	Name string `json:"name"`
	JavaId string `json:"java_id"`
	IsDaemon bool `json:"is_daemon"`
	Prio int `json:"prio"`
	OsPrio int `json:"os_prio"`
	Tid string `json:"tid"`
	Nid string `json:"nid"`
	NativeThreadState string `json:"native_thread_state"`
	ConditionAddress string `json:"condition_address"`
	JavaThreadState string `json:"java_thread_state"`
	JavaStateClarification string `json:"java_state_clarification"`
	StacktraceLines []StacktraceLine `json:"stacktrace_lines"`
}

type StacktraceLine struct {
	Id int `json:"id"`
	Kind string `json:"kind"`
	LineNumber int `json:"line_number"`
	LockAddress *string `json:"lock_address,omitempty"`
	LockedClass *string `json:"locked_class,omitempty"`
	LockClass *string `json:"lock_class,omitempty"`
	JavaClass *string `json:"java_class,omitempty"`
	JavaMethod *string  `json:"java_method,omitempty"`
	SourceLine *string `json:"source_line,omitempty"`
	SourceFile *string `json:"source_file,omitempty"`
}

func listThreads(w http.ResponseWriter, r *http.Request) {
	requestedId := pat.Param(r, "id")

	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT thread.id, thread.name, thread.java_id, thread.is_daemon,
																thread.prio, thread.os_prio, thread.tid, thread.nid,
																thread.native_thread_state, thread.condition_address,
																thread.java_thread_state, thread.java_state_clarification,
																line.id, line.kind, line.line_number, line.lock_address,
																line.locked_class, line.lock_class, line.java_class,
																line.java_method, line.source_line, line.source_file
												 FROM java_threads as thread, stacktrace_lines as line
												 WHERE thread.id = line.java_thread_id
												   AND thread.threaddump_id = ?
												 ORDER BY thread.id, line.line_number`, requestedId)
	if (err != nil) {
		log.Fatal(err)
	}

  threads := []JavaThread{}
	for rows.Next() {
	  var thread JavaThread
		var line StacktraceLine
		err := rows.Scan(&thread.Id, &thread.Name, &thread.JavaId, &thread.IsDaemon,
		                 &thread.Prio, &thread.OsPrio, &thread.Tid, &thread.Nid,
						       	 &thread.NativeThreadState, &thread.ConditionAddress,
						       	 &thread.JavaThreadState, &thread.JavaStateClarification,
						       	 &line.Id, &line.Kind, &line.LineNumber, &line.LockAddress,
						       	 &line.LockedClass, &line.LockClass, &line.JavaClass,
						       	 &line.JavaMethod, &line.SourceLine, &line.SourceFile)
	  if err != nil {
		  log.Fatal(err)
	  }
	  if (len(threads) == 0 || threads[len(threads)-1].Id != thread.Id) {
			threads = append(threads, thread)
		}

		threads[len(threads)-1].StacktraceLines = append(threads[len(threads)-1].StacktraceLines, line)
	}

	b, err := json.Marshal(threads)
	if err != nil {
	  log.Fatal(err)
	}
	w.Write(b)


	/*
	encoder := json.NewEncoder(w)
	err = encoder.Encode(threads)
	*/
}

func listThreaddumps(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, application, host, upload_time FROM threaddumps")
	if err != nil {
		log.Fatal(err)
	}

	dumps := []Threaddump{}
	for rows.Next() {
		var dump Threaddump
		err := rows.Scan(&dump.Id, &dump.Application, &dump.Host, &dump.Uploaded)
		if (err != nil) {
			log.Fatal(err)
		}
		dumps = append(dumps, dump)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(dumps)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/threaddumps"), listThreaddumps)
	mux.HandleFunc(pat.Get("/threads/:id"), listThreads)

  handler := cors.Default().Handler(mux)
	http.ListenAndServe("localhost:8000", handler)
}

func sqliteTest() {
	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int
		application string
		host string
		time time.Time
	)
	defer db.Close()

	rows, err := db.Query("select id, application, host, upload_time from main.threaddumps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &application, &host, &time)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, application, host, time)
	}
	if rows.Err() != nil {
		log.Fatal(err)
	}
}
