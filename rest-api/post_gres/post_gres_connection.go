package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Println("parseURL error")
		log.Fatal(err.Error())
	}

	//`sql.Open()` will create a database instance for us to work with

	//`db` should contain the necessary information for the db object
	//`dbErr` will obviously house the error associated with opening this database
	db, err = sql.Open("postgres", pgUrl)

	if db == nil {
		log.Fatal("db is null fromhere")

	}

	if err != nil {
		log.Println("sqlOpen error")
		log.Fatal(err.Error())
	}

	//`db.Ping()` sends a ping to the database, i presume lol

	//According to the docs: Ping verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	err = db.Ping()

	if err != nil {
		log.Println("Ping error")
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()

	//A route accessed by call the "/books" endpoint with the http method "GET"
	router.HandleFunc("/books", getBooks).Methods("GET")

	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(write http.ResponseWriter, request *http.Request) {
	var book Book = Book{}
	books = []Book{}

	//Query selects all the items in the books database
	if db == nil {
		log.Fatal("Db is null")
	}

	rows, queryErr := db.Query("select * from books")

	if rows == nil {
		log.Fatal("Rows is null")
	}

	if queryErr != nil {
		log.Println("Queury error")
		log.Fatal(queryErr.Error())
	}

	defer rows.Close()

	//rows.Next() should act like one of those java cursor functions and cycles through
	//each row in the database, reading it's values
	for rows.Next() {
		scanError := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if scanError != nil {
			log.Fatal(scanError)
		}

		books = append(books, book)
	}

	json.NewEncoder(write).Encode(books)

}

func getBook(write http.ResponseWriter, request *http.Request) {
	var book Book
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
	var book Book
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
	var book Book
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
