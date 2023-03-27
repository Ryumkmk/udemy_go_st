package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user =test_user dbname = testdb password = password sslmode = disable")
	if err != nil {
		log.Println(err)
	}
	defer Db.Close()

	//C

	/*

		cmd := "INSERT INTO persons (name , age) VALUES ($1,$2)"
		_, err := Db.Exec(cmd, "Ryu", 22)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		//R
		cmd := "SELECT * FROM persons WHERE age = $1"
		row := Db.QueryRow(cmd, 22)
		var p Person
		err = row.Scan(&p.Name, &p.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("NO ROWS")
			} else {
				log.Println(err)
			}
		}
		fmt.Println(p.Name, p.Age)

		cmd = "SELECT * FROM persons"
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		var pp []Person
		for rows.Next() {
			var p Person
			err = rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Println(err)
			}
			pp = append(pp, p)
		}

		err = row.Err()
		if err != nil {
			log.Fatalln(err)
		}
		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}

	*/

	//U

	cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	_, err := Db.Exec(cmd, 23, "Ryu")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		//D

		cmd := "DELETE FROM persons WHERE name = $1"
		_,err = Db.Exec(cmd,"Ryu")
		if err != nil {
			log.Fatalln(err)
		}
	*/
}
