package main

import "fmt"

// It iterates over the array, and for each element, it checks if it's already in the map. If it is, it
// increments the count and removes the element from the map. If it isn't, it adds it to the map
func sockMerchant(n int32, ar []int32) int32 {
	count := 0
	check := make(map[int32]bool)
	for _, num := range ar {

		val := check[num]

		fmt.Println(check)
		if val {
			count++
			check[num] = false
			continue
		}

		check[num] = true
	}
	return int32(count)
}
