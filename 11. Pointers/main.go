package main

import "fmt"

func NameChangeByRef(name *string, otherName string) {
	*name = otherName
}

func main() {
	name := "Daniel"
	fmt.Println("My name is:", name) //Just realize that a space is added for each argument

	NameChangeByRef(&name, "Eli") //Changing name by reference

	fmt.Println("My name is now:", name)
}
