package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/full-cycle-2.0-hexagonal-architecture/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.IProduct, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}