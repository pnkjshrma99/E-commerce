// models/db.go
package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
)

type DBHandler interface {
	ListProducts() ([]Product, error)
	CreateProduct(*Product) error
	// Add more database methods...
}

type Datastore struct {
	DB *sql.DB
}

func ConnectDB() (*Datastore, error) {
	db, err := sql.Open("postgres", "user=postgres password=secret dbname=ecommerce sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Datastore{DB: db}, nil
}

func (db *Datastore) ListProducts() ([]Product, error) {
	rows, err := db.DB.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (db *Datastore) CreateProduct(product *Product) error {
	_, err := db.DB.Exec("INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.Stock)
	return err
}

// Add more database methods for other models...