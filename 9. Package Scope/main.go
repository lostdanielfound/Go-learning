package main

import (
	"fmt"
)

func main() {
	sayHello("Daniel")

	fmt.Println("Reading points from main")
	for _, value := range points {
		fmt.Printf("%v", value)
	}
}
