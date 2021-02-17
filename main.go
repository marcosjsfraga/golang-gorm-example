package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PsqlConn is a connection string
var PsqlConn string = "host=localhost port=5432 user=declabs dbname=gorm_example password=1q2w3e"

func main() {
	fmt.Println("Go ORM Turorial...")

	// InitialMigration()

	handleRequest()
}

func handleRequest() {
	// Create a router
	router := mux.NewRouter().StrictSlash(true)

	// Index
	router.HandleFunc("/", index).Methods("GET")
	// User/Create
	router.HandleFunc("/user", CreateUser).Methods("POST")
	// User/All
	router.HandleFunc("/users", AllUsers).Methods("GET")
	// User/Update
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	// User/Delete
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	// Listen port
	log.Fatal(http.ListenAndServe(":8081", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go-ORM Turorial...")
}
