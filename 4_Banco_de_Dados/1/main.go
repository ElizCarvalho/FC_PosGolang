package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

// Função para criar um novo produto
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

// Método para formatar a string do produto
func (p *Product) String() string {
	return fmt.Sprintf("Produto %v possui o preco de R$ %.2f", p.Name, p.Price)
}

// Método para printar o produto com separador
func (p *Product) PrintWithSeparator() {
	fmt.Println(p.String())
	fmt.Println("--------------------")
}

// Função para printar a lista de produtos
func printProducts(products []Product) {
	for _, p := range products {
		fmt.Println(p.String())
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("TV", 2529.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.PrintWithSeparator()

	product.Price = 879.90
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.PrintWithSeparator()

	product, err = selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	product.PrintWithSeparator()

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	printProducts(products)

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

// Função para inserir um produto
func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

// Função para atualizar um produto
func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
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

// Função para selecionar um produto
func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Função para selecionar todos os produtos
func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
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

// Função para deletar um produto
func deleteProduct(db *sql.DB, id string) error {
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
