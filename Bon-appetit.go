package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here

	var suma int32

	for _, sum := range bill {
		suma += sum

	}
	total := (suma - bill[k]) / 2

	if total == b {
		fmt.Println("Bon Appetit")
	} else {

		total = b - total
		fmt.Println(total)
	}

}
