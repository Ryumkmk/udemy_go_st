package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}
func main() {
	user1 := User{Name: "Ryu"}
	user1.SayName()
	user1.SetName("Ko")
	user1.SayName()

}
