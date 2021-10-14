package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	query := "create table products(id string, name string, price float, status string);"

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	query := `insert into products values("abc", "Product Test", 0, "disabled");`

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}
