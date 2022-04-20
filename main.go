package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/*const (
	addr := os.Getenv("address")
	port := os.Getenv("port")
	user := os.Getenv("user")
	pass := os.Getenv("password")
	dbname := os.Getenv("db")
)*/

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connect()
}

func connect() {
	addr := os.Getenv("address")
	port := os.Getenv("port")
	user := os.Getenv("user")
	pass := os.Getenv("password")
	dbname := os.Getenv("db")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", addr, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	rows, err := db.Query(`SELECT "nimi", "ika" FROM "henkilot"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		CheckError(err)

		fmt.Println(name, roll)
	}

	defer db.Close()

	fmt.Println("Successfully connected!")
}

// func add_person() {}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
