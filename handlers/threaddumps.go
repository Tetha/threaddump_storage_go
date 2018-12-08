package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var threaddumpsTemplates = template.Must(template.ParseGlob("templates/*.html"))

// ListThreaddumps lists all threaddumps in the system
func (env *RuntimeEnvironment) ListThreaddumps(w http.ResponseWriter, r *http.Request) {
	dumps, err := env.db.ListAllThreaddumps()
	if err != nil {
		log.Printf("Error getting dumps from db: %s", err)
		http.Error(w, "500: Error getting dumps from database", 500)
	}
	err = templates.ExecuteTemplate(w, "threaddumps_list.html", dumps)
	if err != nil {
		http.Error(w, "500: Error rending template", 500)
		log.Printf("Error rendering template: %s", err)
		return
	}
}
