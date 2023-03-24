package main

import "fmt"

type Stringfy interface {
	Tostring() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Tostring() string {
	return fmt.Sprintf("Name=%v,Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) Tostring() string {
	return fmt.Sprintf("Number=%v,Model=%v", c.Number, c.Model)

}
func main() {
	vs := []Stringfy{
		&Person{"Ryu", 10},
		&Car{"123-33", "Doge"},
	}

	for _, v := range vs {
		fmt.Println(v.Tostring())
	}
}
