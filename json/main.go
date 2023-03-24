package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id,string"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A         `json:"A"`
}

func main() {

	u := new(User)
	u.ID = 123
	u.Name = "Ryu"
	u.Email = "ryu@gmail.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	u2 := new(User)
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*u2)
}
