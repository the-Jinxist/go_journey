package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDriver() *sql.DB {

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

	return db
}
