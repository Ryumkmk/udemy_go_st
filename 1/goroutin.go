package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Printf("Sub loop \n")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	go sub()
	go sub()
	
	for {
		fmt.Printf("Main loop \n")
		time.Sleep(200 * time.Millisecond)
	}
}
