package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, *product)

	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product Product) error {

	// os "?" protegem o banco de SQL injection
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	// Aqui os "?" serão substituídos pelos dados da struct
	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}
