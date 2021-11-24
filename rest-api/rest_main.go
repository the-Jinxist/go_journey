package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var bools []Book

func main() {
	router := mux.NewRouter()

	//Mux library helps Golang developers to create routes

	//A route accessed by call the "/books" endpoint with the http method "GET"
	router.HandleFunc("/books", getBooks).Methods("GET")

	router.HandleFunc("/books{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(write http.ResponseWriter, request *http.Request) {
	log.Println("getBooks()")
}

func getBook(write http.ResponseWriter, request *http.Request) {
	log.Println("getBook()")
}

func addBook(write http.ResponseWriter, request *http.Request) {
	log.Println("addBook()")
}

func updateBook(write http.ResponseWriter, request *http.Request) {
	log.Println("updateBook()")
}

func removeBook(write http.ResponseWriter, request *http.Request) {
	log.Println("removeBook()")
}
