package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

	product.Price = 100.0

	err = updateProduct(db, *product)
	if err != nil {
		panic(err)
	}

	// Contexto para o select
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Selecionar um produto
	// p, err := selectProduct(ctx, db, product.ID)

	products, err := selectAllProducts(ctx, db)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(ctx, db, product.ID)

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

func updateProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("UPDATE products set name = ?, price = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	// QueryRow seleciona apenas uma ocorrência
	// Scan joga o resultado para alguma variável (No caso no endereço de memória do objeto que criamos)
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(ctx context.Context, db *sql.DB) ([]Product, error) {

	rows, err := db.QueryContext(ctx, "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	// Percorre linha a linha os produtos
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil

}

func deleteProduct(ctx context.Context, db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
