package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	db sql.DB
}

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
	if err != nil {
		panic(err)
	}

	//db.get_person()

	rows, err := db.Query(`SELECT "nimi", "ika" FROM "henkilot"`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var age int

		err = rows.Scan(&name, &age)
		if err != nil {
			panic(err)
		}

		fmt.Println(name, age)

		defer db.Close()

		fmt.Println("Successfully connected!")
	}

	/*func (db *DB) get_person() {

		rows, err := db.db.Query(`SELECT "nimi", "ika" FROM "henkilot"`)
		CheckError(err)

		defer rows.Close()
		for rows.Next() {
			var name string
			var age int

			err = rows.Scan(&name, &age)
			CheckError(err)

			fmt.Println(name, age)
		}
	}*/

	// func add_person() {}

}
