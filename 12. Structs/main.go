package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin) //Use bufio to reader from stdin from the os
	fmt.Print("Create a new bill name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	name = strings.TrimSpace(name)
	b := NewBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	mybill := NewBill("Mario")
	AddToBill(mybill, "Pizza", 5.15)
	AddToBill(mybill, "Burger", 10.25)
	AddToBill(mybill, "Soda", 1.25)
	PrintBill(mybill)

	lolBill := CreateBill()
	PrintBill(lolBill)
}
