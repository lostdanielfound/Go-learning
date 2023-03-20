package main //Always need to have this

import (
	"fmt"
)

var points = []int{1, 2, 3, 4} //points is declared in here, but since we included package main, we can also use this array within the main.go

func sayHello(name string) { //Says hello to the name that gets passed in
	fmt.Println("Hello ", name)
	fmt.Printf("These are the values of the array: %v\n", points)
}
