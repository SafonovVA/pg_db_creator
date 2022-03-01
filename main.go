package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	fmt.Print("Enter name: ")
	var name string
	_, err = fmt.Scan(&name)
	if err != nil {
		panic(err)
	}
	dbName := name + "_db"
	userName := name + "_admin"

	_, err = db.Exec("CREATE DATABASE " + dbName + ";")
	if err != nil {
		panic(err)
	}
	fmt.Println("database " + dbName + " created")
	_, err = db.Exec("CREATE USER " + userName + " WITH ENCRYPTED PASSWORD '1'")
	if err != nil {
		panic(err)
	}
	fmt.Println("user " + userName + " created")
	_, err = db.Exec("GRANT ALL PRIVILEGES ON DATABASE " + dbName + " TO " + userName + ";")
	if err != nil {
		panic(err)
	}
	fmt.Println("privileges on " + dbName + " granted to " + userName)
}
