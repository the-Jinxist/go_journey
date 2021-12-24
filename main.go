package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go_journey/controllers"
	"go_journey/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDriver()
	booksController := controllers.Controller{}

	router := mux.NewRouter()

	//A route accessed by call the "/books" endpoint with the http method "GET"
	router.HandleFunc("/books", booksController.GetBooks(db)).Methods("GET")

	router.HandleFunc("/books/{id}", booksController.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", booksController.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", booksController.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", booksController.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
