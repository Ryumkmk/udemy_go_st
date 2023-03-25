package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS person(
	// 	name STRING,
	// 	age INT
	// )`

	// cmd := "INSERT INTO person (name, age) VALUES(?, ?) "

	// cmd := "UPDATE person SET age = ? WHERE name = ?"

	// cmd := "SELECT * FROM person WHERE age =?"

	cmd := "SELECT * FROM person "

	// row := Db.QueryRow(cmd, 19)

	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	// err := row.Scan(&p.Name, &p.Age)
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// _, err := Db.Exec(cmd, "hanako", 19)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			log.Println("No row")
	// 		} else {
	// 			log.Println(err)
	// 		}
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(p.Name, p.Age)
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
