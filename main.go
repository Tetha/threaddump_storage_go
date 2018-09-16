package main

import (
	"log"
	"time"

	"database/sql"

	"html/template"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tetha/threaddumpstorage-go/database"
	"github.com/tetha/threaddumpstorage-go/handlers"
	"github.com/tetha/threaddumpstorage-go/upload"

	"github.com/gorilla/mux"

	"net/http"
	"net/http/pprof"
)

type Threaddump struct {
	ID          int       `json:"id"`
	Application string    `json:"application"`
	Host        string    `json:"host"`
	Uploaded    time.Time `json:"uploaded"`
}

type JavaThread struct {
	Id                     int              `json:"id"`
	Name                   string           `json:"name"`
	JavaId                 string           `json:"java_id"`
	IsDaemon               bool             `json:"is_daemon"`
	Prio                   int              `json:"prio"`
	OsPrio                 int              `json:"os_prio"`
	Tid                    string           `json:"tid"`
	Nid                    string           `json:"nid"`
	NativeThreadState      string           `json:"native_thread_state"`
	ConditionAddress       string           `json:"condition_address"`
	JavaThreadState        string           `json:"java_thread_state"`
	JavaStateClarification string           `json:"java_state_clarification"`
	StacktraceLines        []StacktraceLine `json:"stacktrace_lines"`
}

type StacktraceLine struct {
	Id          int     `json:"id"`
	Kind        string  `json:"kind"`
	LineNumber  int     `json:"line_number"`
	LockAddress *string `json:"lock_address,omitempty"`
	LockedClass *string `json:"locked_class,omitempty"`
	LockClass   *string `json:"lock_class,omitempty"`
	JavaClass   *string `json:"java_class,omitempty"`
	JavaMethod  *string `json:"java_method,omitempty"`
	SourceLine  *string `json:"source_line,omitempty"`
	SourceFile  *string `json:"source_file,omitempty"`
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := mux.NewRouter()

	// ----------------
	// Setup pprof http
	// ----------------
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// --------------
	// Static Content
	// --------------
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))

	db, err := database.NewSQLiteStore("./threaddump.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	env := handlers.NewEnvironment(db)
	// -------------------------
	// Threaddump detail section
	// -------------------------
	s := r.PathPrefix("/threaddump/{dumpId:[0-9]+}").Subrouter()
	s.HandleFunc("/pools", env.ListThreadpools)
	s.HandleFunc("/threads", env.ListThreads)

	// -----------------------------
	// Threaddump management section
	// -----------------------------
	r.HandleFunc("/upload", upload.HandleUpload)
	r.HandleFunc("/threaddumps", listThreaddumps)

	http.Handle("/", r)

	log.Print("Serving on 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
func listThreads(w http.ResponseWriter, r *http.Request) {
	// TODO: validation

	r.ParseForm()
	requestedID := r.URL.Path[len("/threads/"):]
	if requestedID == "" {
		http.Error(w, "400: Missing threaddump id", 400)
		return
	}

	from := r.FormValue("from")
	if from == "" {
		from = "0"
	}
	to := r.FormValue("to")
	if to == "" {
		to = "20"
	}
	log.Printf("Listing from %s to %s", from, to)

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
												 ORDER BY thread.id, line.line_number
												 LIMIT ? OFFSET ?`, requestedID, to, from)
	if err != nil {
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
		if len(threads) == 0 || threads[len(threads)-1].Id != thread.Id {
			threads = append(threads, thread)
		}

		threads[len(threads)-1].StacktraceLines = append(threads[len(threads)-1].StacktraceLines, line)
	}
	err = templates.ExecuteTemplate(w, "threads_list.html", threads)
	if err != nil {
		http.Error(w, "500: Error rending template", 500)
		log.Printf("Error rendering template: %s", err)
		return
	}
}
*/
func listThreaddumps(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, application, host, upload_time FROM threaddumps")
	if err != nil {
		log.Printf("Error with db query: %s", err)
		http.Error(w, "500: Database Error", 500)
	}

	dumps := []Threaddump{}
	for rows.Next() {
		var dump Threaddump
		err := rows.Scan(&dump.ID, &dump.Application, &dump.Host, &dump.Uploaded)
		if err != nil {
			log.Printf("Error scanning result rows: %s", err)
			http.Error(w, "500: Database Error", 500)
		}
		dumps = append(dumps, dump)
	}
	err = templates.ExecuteTemplate(w, "threaddumps_list.html", dumps)
	if err != nil {
		http.Error(w, "500: Error rending template", 500)
		log.Printf("Error rendering template: %s", err)
		return
	}
}

func sqliteTest() {
	db, err := sql.Open("sqlite3", "./threaddump.db")
	if err != nil {
		log.Fatal(err)
	}

	var (
		id          int
		application string
		host        string
		time        time.Time
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
