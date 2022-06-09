package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var totalApples int32
	var totalOranges int32

	for _, value := range apples {
		if a+value >= s && a+value <= t {
			totalApples++
		}
	}

	for _, value := range oranges {
		if b+value >= s && b+value <= t {
			totalOranges++
		}
	}

	fmt.Println(totalApples)
	fmt.Println(totalOranges)

}
