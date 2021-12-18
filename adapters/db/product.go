package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/full-cycle-2.0-hexagonal-architecture/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func CreateTable(db *sql.DB) {
	query := "create table if not exists products(id string, name string, price float, status string);"

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
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

func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	var rows int

	p.db.QueryRow("Select count(*) from products where id=?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.db.Exec(
		"update products set name = ?, price=?, status=? where id = ?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetId(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
