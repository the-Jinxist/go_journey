package bookRepository

import (
	"database/sql"
	"go_journey/models"
	"log"
)

type BookRepository struct{}

func (b BookRepository) GetBooksFromDB(db *sql.DB) (books []models.Book, error error) {
	book := models.Book{}

	rows, queryErr := db.Query("select * from books")

	if rows == nil {
		log.Fatal("Rows is null")

		books = []models.Book{}
		return books, nil
	}

	if queryErr != nil {
		log.Fatal(queryErr.Error())
		error = queryErr
		return nil, error
	}

	defer rows.Close()

	//rows.Next() should act like one of those java cursor functions and cycles through
	//each row in the database, reading it's values
	for rows.Next() {
		scanError := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if scanError != nil {
			return nil, scanError
		}

		books = append(books, book)
	}

	return books, error
}
