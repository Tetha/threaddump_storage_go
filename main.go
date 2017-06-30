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
