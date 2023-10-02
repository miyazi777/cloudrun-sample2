package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "test_db"
	USER     = "user"
	PASSWORD = "pass"
)

func main() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name FROM sample")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d name: %s\n", id, name)
	}
}
