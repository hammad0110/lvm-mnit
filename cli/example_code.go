package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	.Db.exec()
	// .exec

	var err error
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", searchHandler)
	http.ListenAndServe(":8080", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	// Potential SQL Injection
	rows, err := db.Query("SELECT * FROM users WHERE username='" + query + "'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "User ID: %d, Username: %s\n", id, username)
	}
}
