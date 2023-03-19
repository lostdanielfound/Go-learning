package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	var ages [3]int = [3]int{1, 2, 3} //One way of declaring an array using the var keyword
	amountOfBoys := [3]int{3, 3, 2}   //another way of declaring an array using the := symbol

	fmt.Println("Contents of the ages array[]:", ages)
	fmt.Println("Contents of the amountOfBoys array[]:", amountOfBoys)

	numbers := []int{}           //Creating a slice with no elements initally
	numbers = append(numbers, 1) //Appending 1 element
	numbers = append(numbers, 2)
	numbers = append(numbers, 120)
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Println("Last element within numbers: ", numbers[len(numbers)-1])
}
