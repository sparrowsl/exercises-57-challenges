package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type application struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("sqlite", "./breve.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to db successfully")

	app := &application{
		db: db,
	}

	serv := http.Server{
		Handler:     app.routes(),
		Addr:        ":5000",
		ReadTimeout: time.Second * 5,
	}

	serv.ListenAndServe()
}