package main

import (
	"fmt"
)

func main() {
	homes := map[int]string{
		111: "Mario res.",
		222: "Luigi res.",
		333: "Eli res.",
		444: "Daniel res.",
	}

	fmt.Println("Who lives at res 444: ", homes[444])
	if homes[222] != "Mario res." {
		fmt.Println("Mario does not live at 222")
	} else {
		fmt.Println("Mario Does live at 222! Get him bruh!!!")
	}

	for key, value := range homes {
		fmt.Printf("%s lives at %v\n", value, key)
	}

	//Moving mario to a different res
	homes[555] = "Mario res." //moving mario to a different key
	homes[111] = ""           //creating a Vacent value

	fmt.Printf("%s now lives at 555", homes[555])
}
