package main

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.println(u.Name)
}
func main() {
	user1 := User{Name: "Ryu"}
	user1.SayName()
}
