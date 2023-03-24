package main

import "fmt"

type User struct {
	Name string
	Age  int
}
type Users []*User

func main() {
	user1 := User{"user1", 10}
	user2 := User{"user2", 20}
	user3 := User{"user3", 30}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)

	}

}
