package dbconnect

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("postgres", "host="+"localhost"+" user="+"postgres"+" dbname="+"urlshorten"+" sslmode=disable password="+"sumeet")

	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//fmt.Println("Connected")

	return db
}
