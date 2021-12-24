package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go_journey/controllers"
	"go_journey/driver"
	"go_journey/models"

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

	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBook(write http.ResponseWriter, request *http.Request) {
	var book models.Book
	params := mux.Vars(request)

	row := db.QueryRow("select * from books where id=$1", params["id"])

	if row == nil {
		json.NewEncoder(write).Encode("No book with this id exists")
	} else {
		row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		json.NewEncoder(write).Encode(book)

	}

}

func addBook(write http.ResponseWriter, request *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(request.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(write).Encode(bookID)

}

func updateBook(write http.ResponseWriter, request *http.Request) {
	var book models.Book
	json.NewDecoder(request.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id;", book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		log.Fatal(err)
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(write).Encode(rowsUpdated)

}

func removeBook(write http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	result, err := db.Exec("delete from books where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(write).Encode(rowsUpdated)

}
