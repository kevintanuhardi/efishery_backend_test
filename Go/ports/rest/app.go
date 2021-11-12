package rest

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"phone" TEXT,
		"name" TEXT,
		"rule" TEXT,
		"password" TEXT(4)
	  );`

	statement, err := db.Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err)
		return err
	}
	statement.Exec() // Execute SQL Statements
	return nil
}

// Application for rest service
func Application(cfg *Config) error {
	db, err := AppWithGorm(cfg)
	if err != nil {
		return err
	}

	server := Server{cfg: cfg.Cfg, router: chi.NewRouter(), db: db}
	return server.Run()
}

func AppWithGorm(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		return nil, err
	}
	defer db.Close() // Defer Closing the database
	err = createTable(db) // Create Database Tables
	return db, err
}
