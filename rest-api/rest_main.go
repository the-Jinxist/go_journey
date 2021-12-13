package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Report interface {
	Dedication()
	Summary()
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func (b *Book) changeBookName(name string) {
	b.Title = name
}

func main() {
	//Mux library helps Golang developers to create routes
	router := mux.NewRouter()

	book := Book{}
	book.changeBookName("Favour")

	log.Println(book)

	books = append(books, Book{ID: 1, Title: "GoLang Stuff", Author: "Favour", Year: "2021"},
		Book{ID: 2, Title: "GoLang Stuff Pt. 2", Author: "Favour", Year: "2021"},
		Book{ID: 3, Title: "GoLang Stuff Pt. 3", Author: "Favour", Year: "2021"},
		Book{ID: 4, Title: "GoLang Stuff Pt. 4", Author: "Favour", Year: "2021"},
		Book{ID: 5, Title: "GoLang Stuff Pt. 5", Author: "Favour", Year: "2021"})

	//A route accessed by call the "/books" endpoint with the http method "GET"
	router.HandleFunc("/books", getBooks).Methods("GET")

	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(write http.ResponseWriter, request *http.Request) {
	//Seems json encode writes the books to the response writer
	json.NewEncoder(write).Encode(books)
}

func getBook(write http.ResponseWriter, request *http.Request) {
	log.Println("getBook()")

	//`mux.Vars(http.Request)` is used to get the request parameters in a request, I guess
	params := mux.Vars(request)
	intID, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Error with converting string `id` to int")
	}

	for _, book := range books {
		if book.ID == intID {
			json.NewEncoder(write).Encode(&book)
		}
	}
	log.Println(params)
}

func addBook(write http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(write).Encode(books)
}

func updateBook(write http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(write).Encode(books)
}

func removeBook(write http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(write).Encode(books)
}
