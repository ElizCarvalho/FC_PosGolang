package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// GetDB retorna uma conexão com o banco de dados
func GetDB() *sql.DB {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./database.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Criar tabelas se não existirem
	if err := createTables(db); err != nil {
		log.Fatalf("Erro ao criar tabelas: %v", err)
	}

	return db
}

// createTables cria as tabelas necessárias no banco de dados
func createTables(db *sql.DB) error {
	categoryTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		description TEXT
	);`

	if _, err := db.Exec(categoryTable); err != nil {
		return err
	}

	return nil
}
