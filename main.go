package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"

	"forum-db-interface/controller"
	dbpkg "forum-db-interface/db"
)

func main() {
	db, err := dbpkg.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	c := controller.New(db)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      c,
	}

	log.Printf("Server successfully started at port %v\n", server.Addr)
	log.Println(server.ListenAndServe())
}
