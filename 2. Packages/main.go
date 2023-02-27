package main //This is a comment. every go program is made up of packages. 
//This program has the main package

//You may import packages like this:
//import (
	//"fmt"
	//"math/rand"
//)

//Or like this:
import "fmt"
import "math/rand"
import "math"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("This is a random number using the rand package: ", rand.Intn(10)) //I believe this rand.Intn function produces a number between 0 to 10 including 10 
	fmt.Println("This is the square root of 4: ", math.Sqrt(4)) //NOTICE THIS THESE NAMES .Sqrt and Intn are capitalized. You must refer to exported object with a Captial letter. 
	fmt.Println("This is the square root of 13: ", math.Sqrt(13))
	fmt.Println("This is the sum of 3 and 5: ", add(3, 5))
}