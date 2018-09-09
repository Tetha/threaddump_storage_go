package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/tetha/threaddumpstorage-go/model"

	"html/template"

	"github.com/gorilla/mux"
)

type threadPoolDetection struct {
	ThreadPools    map[string][]model.JavaThreadHeader
	UnknownThreads []model.JavaThreadHeader
}

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

	detectedPools := figureOutThreadpools(threads)

	err = templates.ExecuteTemplate(w, "threadpool_list.html", detectedPools)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
	}

}

func figureOutThreadpools(threads []model.JavaThreadHeader) (pools threadPoolDetection) {
	pools.ThreadPools = make(map[string][]model.JavaThreadHeader)
	pools.UnknownThreads = threads

	findElasticsearchThreadpools(&pools)
	findNumberedThreads(&pools)
	return
}

func findNumberedThreads(pools *threadPoolDetection) {
	var newUnknownThreads []model.JavaThreadHeader
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
	var newUnknownThreads []model.JavaThreadHeader
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
