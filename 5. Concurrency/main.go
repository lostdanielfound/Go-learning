package main

import "fmt"

func sum(c chan int) {
	x := <- c  //data from channel is assigned to x first 
	y := <- c  //data from channel is assigned to y next
	c <- x + y //the sum of x and y is then sent back to the channel to be "return"
}

func main() {
	c := make(chan int) //make c a channel
	go sum(c)
	c <- 10 //throw 10 into the channel 
	c <- 15 //throw 15 into the channel 
	r := <- c //Whatever is in the channel will be assigned to r
	fmt.Println("This is the value of r: ", r)
}