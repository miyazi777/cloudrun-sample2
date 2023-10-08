package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, database)
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{ID: id, Name: name})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
