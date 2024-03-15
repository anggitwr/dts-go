package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Person struct {
	Id      int
	Name    string
	Address string
}

func main() {
	fmt.Println("hello")

	connectionString := "host=localhost port=5432 user=postgres password=postgres dbname=hacktiv sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	getAll(db)

	insert(db, "harto", "bandung")

}

func insert(db *sql.DB, name, address string) {
	query := "insert into person(name, address) values($1, $2)"
	var persons = []Person{}
	rows, err := db.Query(query, name, address)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var p = Person{}

		err = rows.Scan(&p.Id, &p.Name, &p.Address)
		if err != nil {
			continue
		}
		persons = append(persons, p)
	}

	fmt.Println(persons)
}

func getAll(db *sql.DB) {
	query := "select * from person"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var persons = []Person{}

	for rows.Next() {
		var p = Person{}

		err = rows.Scan(&p.Id, &p.Name, &p.Address)
		if err != nil {
			continue
		}
		persons = append(persons, p)
	}

	fmt.Println(persons)
}
