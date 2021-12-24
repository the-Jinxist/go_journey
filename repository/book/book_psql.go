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

func (b BookRepository) GetBookFromDB(db *sql.DB, id string) (book models.Book, error error) {
	row := db.QueryRow("select * from books where id=$1", id)

	if row == nil {
		return models.Book{}, error
	} else {
		scanError := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if scanError != nil {
			return book, scanError
		}
		return book, nil

	}
}

func (b BookRepository) AddBookToDB(db *sql.DB, book models.Book) (bookID int, error error) {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	if err != nil {
		return 0, err
	} else {
		return bookID, nil
	}
}

func (b BookRepository) UpdateBookOnDB(db *sql.DB, book models.Book) (bookID int, error error) {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id;", book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), nil
}

func (b BookRepository) DeleteBooks(db *sql.DB, id string) error {

	result, err := db.Exec("delete from books where id = $1", id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
