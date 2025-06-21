package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

func PostgreSqlInit() *DataBase {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка соединения с БД: %v", err)
	}
	createdb := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, 
		name TEXT,
		email TEXT
	);
	`
	_, err = db.Exec(createdb)
	if err != nil {
		log.Fatal(err)
	}

	return &DataBase{
		DB: db,
	}
}
