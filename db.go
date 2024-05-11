package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/cetec")
	if err != nil {
		return nil, err
	}
	return db, nil
}
