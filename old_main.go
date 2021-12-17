// package main

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/full-cycle-2.0-hexagonal-architecture/adapters/db"
// 	"github.com/full-cycle-2.0-hexagonal-architecture/application"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	Db, err := sql.Open("sqlite3", ":memory:")

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	db.CreateTable(Db)

// 	productDbAdapter := db.NewProductDb(Db)
// 	productService := application.NewProductService(productDbAdapter)

// 	product, _ := productService.Create("Example Product", 10)

// 	log.Print(productService.Get(product.GetId()))
// }
