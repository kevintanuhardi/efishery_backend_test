package rest

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTable(db *gorm.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"phone" TEXT,
		"name" TEXT,
		"role" TEXT,
		"password" TEXT(4)
	  );`
	db.Exec(createTableSQL) // Execute SQL Statements
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

func AppWithGorm(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./efishery-database.db"), &gorm.Config{})
	// db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		return nil, err
	}
	err = createTable(db) // Create Database Tables
	return db, err
}
