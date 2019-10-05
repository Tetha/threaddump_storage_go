package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tetha/threaddumpstorage-go/database"
	"github.com/tetha/threaddumpstorage-go/handlers"

	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"net/http"
	"net/http/pprof"
)

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
	r.HandleFunc("/upload", env.HandleUpload)
	r.HandleFunc("/threaddumps", env.ListThreaddumps)

	a := r.PathPrefix("/api").Subrouter()
	a.HandleFunc("/threaddumps", env.ListThreaddumpsAPI)

	aSpecific := a.PathPrefix("/threaddump/{dumpId:[0-9]+}").Subrouter()
	aSpecific.HandleFunc("/pools", env.ListThreadpoolsAPI)
	http.Handle("/", r)

	headersOk := gorilla_handlers.AllowedHeaders([]string{})
	originsOk := gorilla_handlers.AllowedOrigins([]string{"http://127.0.0.1:8080", "http://localhost:8080"})
	methodsOk := gorilla_handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Print("Serving on 8000...")
	log.Fatal(http.ListenAndServe(":8000", gorilla_handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
