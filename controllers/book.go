package controllers

import (
	"database/sql"
	"go_journey/models"
	bookRepository "go_journey/repository/book"
	"go_journey/utils"
	"net/http"
)

type Controller struct{}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {

		booksRepository := bookRepository.BookRepository{}
		books, err := booksRepository.GetBooksFromDB(db)

		if err != nil {
			error := models.Error{Message: err.Error()}
			utils.SendError(write, error, 400)

		} else {
			utils.SendSuccess(write, books, 200)
		}

	}
}
