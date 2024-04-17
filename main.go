package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haji-saklain/demo-api/api"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:Saklain@1@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/hello/{message}", api.HelloHandler(db)).Methods("GET")
	r.HandleFunc("/goodbye/{message}", api.GoodbyeHandler(db)).Methods("GET")
	r.HandleFunc("/count/{count}", api.CountNumbersHandler(db)).Methods("GET")
	r.HandleFunc("/time", api.PrintCurrentTimeHandler(db)).Methods("GET")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
