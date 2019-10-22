package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func homeAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API!")
}

func handleRequest() {
	var PORT = ":8082"
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api", homeAPI).Methods("GET")
	myRouter.HandleFunc("/api/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/api/user/{id}", UserbyID).Methods("GET")
	myRouter.HandleFunc("/api/user", NewUser).Methods("POST")
	myRouter.HandleFunc("/api/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/user/{id}", UpdateUser).Methods("PATCH")
	log.Fatal(http.ListenAndServe(PORT,myRouter))
}

func main() {
	fmt.Println("GO lang ORM Test")
	handleRequest()
}
