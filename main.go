package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequest() {
	//var PORT := 8082
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/add", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/delete", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/update", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082",myRouter))
}

func main() {
	fmt.Println("GO lang ORM Test")
	handleRequest()
}