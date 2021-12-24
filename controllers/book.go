package controllers

import (
	"database/sql"
	"encoding/json"
	"go_journey/models"
	bookRepository "go_journey/repository/book"
	"go_journey/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

var booksRepository = bookRepository.BookRepository{}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {

		books, err := booksRepository.GetBooksFromDB(db)
		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, http.StatusOK)
			return

		} else {
			utils.SendSuccess(write, books, http.StatusOK)

			return
		}

	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {

		params := mux.Vars(request)
		book, err := booksRepository.GetBookFromDB(db, params["id"])

		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, http.StatusNotFound)
			return

		} else {
			utils.SendSuccess(write, book, http.StatusOK)

			return
		}

	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)

		_, err := booksRepository.AddBookToDB(db, book)

		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, http.StatusNotFound)
			return

		} else {
			utils.SendSuccess(write, map[string]string{"message": "New book successfully added"}, http.StatusOK)
			return
		}
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)

		_, err := booksRepository.UpdateBookOnDB(db, book)

		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, http.StatusNotFound)
			return

		} else {
			utils.SendSuccess(write, map[string]string{"message": "Book successfully updated"}, http.StatusOK)
			return
		}
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {

		params := mux.Vars(request)
		id := params["id"]

		err := booksRepository.DeleteBooks(db, id)

		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, http.StatusNotFound)
			return

		} else {
			utils.SendSuccess(write, map[string]string{"message": "Book successfully deleted"}, http.StatusOK)
			return
		}

	}
}
