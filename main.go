package main

import (
	"log"
	"time"

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
	r.HandleFunc("/threaddumps", env.ListThreaddumps)

	http.Handle("/", r)

	log.Print("Serving on 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
