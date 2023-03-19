package main

import (
	"fmt"
)

func main() {
	count := 5
	for i := 0; i < count; i++ {
		fmt.Printf("Current Index is: %v\n", i)
	}

	keyWords := [4]string{"Bruh", "Hella", "lol", "test"}
	for index, value := range keyWords { //The underscore is used for the index, but if we want to omit the usage we use an underscore, same with the value
		fmt.Printf("This is the word in the array at index %v: %v\n", index, value)
	}
}
