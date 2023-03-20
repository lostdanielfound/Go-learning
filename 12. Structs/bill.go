package main

import "fmt"

type bill struct { //struct of a bill object
	name  string             //name of the person whom which the bill is addressed to
	items map[string]float64 //A map of every item ordered and their names followed by their dollor amount
	tip   float64            //An optional tip value for the bill
}

// Make new bills
func NewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.0,
	}

	return b
}

func PrintBill(b bill) {
	fmt.Println("Requestor:", b.name)
	fmt.Println("Billing:")
	for food, price := range b.items {
		fmt.Printf("\t%v - %v\n", food, price)
	}
	fmt.Println("Tip:", b.tip)
}

func AddToBill(b bill, item string, cost float64) {
	//Failure point, if item is already there we should add to the price
	b.items[item] = cost
}
