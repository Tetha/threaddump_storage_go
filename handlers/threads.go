package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/tetha/threaddumpstorage-go/model"

	"github.com/gorilla/mux"
)

var threadTemplates = template.Must(template.ParseGlob("templates/*.html"))

type pagedThreadList struct {
	ThreaddumpID string
	From         int
	Limit        int

	Threads []model.JavaThreadHeader
}

func (list pagedThreadList) PreviousPageStart() int {
	return list.From - list.Limit
}

func (list pagedThreadList) NextPageStart() int {
	return list.From + list.Limit
}

// ListThreads displays a paged list of the threads in a threaddump
func (env *RuntimeEnvironment) ListThreads(w http.ResponseWriter, r *http.Request) {
	// TODO: validation

	threaddumpID := mux.Vars(r)["dumpId"]
	r.ParseForm()

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

	threads, err := env.db.ListPagedThreadHeaders(threaddumpID, limit, from)
	if err != nil {
		log.Print("Error while getting threads", err)
		http.Error(w, "500: database error", 500)
		return
	}
	templateInput := pagedThreadList{threaddumpID, from, limit, threads}
	err = threadTemplates.ExecuteTemplate(w, "threads_list.html", templateInput)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
	}
}
