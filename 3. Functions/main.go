package main

import (
	"fmt"
	"math/rand"
)

func swap(x, y string) (string, string) {
	return y, x //return the two strings that were passed in but in swappped
}

// Function that demostraites naked return with two return values
func tworandInt(num int) (x, y int) {
	x = rand.Intn(num)
	y = rand.Intn(num)
	return
}

func sayString(word string) {
	fmt.Println(word)
}

func cycleNames(n []string, f func(string)) { //A function that takes in an array of strings and uses another function to inake on the array
	for _, word := range n {
		f(word)
	}
}

func main() {
	a, b := swap("world", "Hello,")
	c, d := tworandInt(10)
	fmt.Println("Bruh watch me swap world and hello: ", a, b)
	fmt.Println("These are two random numbers between 0-10: ", c, d)

	someWords := []string{"test1", "test2", "test3"}

	cycleNames(someWords, sayString)
}
