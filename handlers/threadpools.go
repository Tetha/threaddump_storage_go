package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
	"github.com/tetha/threaddumpstorage-go/analysis"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// ListThreadpools is an http handler which displays all detected threadpools
func (env *RuntimeEnvironment) ListThreadpools(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	threaddumpID := mux.Vars(r)["dumpId"]
	if threaddumpID == "" {
		http.Error(w, "400: Missing threaddump id", 400)
		return
	}

	threads, err := env.db.ListAllThreadHeadersInDump(threaddumpID)
	if err != nil {
		http.Error(w, "500: cannot get threads from db", 500)
		return
	}

	detectedPools := analysis.FigureOutThreadpools(threads)

	err = templates.ExecuteTemplate(w, "threadpool_list.html", detectedPools)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
	}

}

func (env *RuntimeEnvironment) ListThreadpoolsAPI(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	threaddumpID := mux.Vars(r)["dumpId"]
	if threaddumpID == "" {
		http.Error(w, "400: Missing threaddump id", 400)
		return
	}

	threads, err := env.db.ListAllThreadHeadersInDump(threaddumpID)
	if err != nil {
		http.Error(w, "500: cannot get threads from db", 500)
		return
	}

	detectedPools := analysis.FigureOutThreadpools(threads)

	buf, err := json.Marshal(detectedPools)
	if err != nil {
		log.Printf("Unable to render json: %s", err)
		http.Error(w, "500: Unable to render json", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(buf)
}
