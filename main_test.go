package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	addr := "localhost"
	port := "5432"
	user := "user"
	pass := "password"
	dbname := "db"

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", addr, port, user, pass, dbname)

	_, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
}
