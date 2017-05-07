package main

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"net/http"
	"goji.io"
	"goji.io/pat"
)

func main() {
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
