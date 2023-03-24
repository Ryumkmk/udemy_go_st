package main

import "fmt"

type T struct {
	User User
	// User
}
type User struct {
	Name string
	Age  int
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	
	t := T{User: User{Name: "Ryu"}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	// fmt.Println(t.Name)
	t.User.SetName("Ko")
}
