package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var contPos, contNeg, contZero float64

	size := float64(len(arr))

	for _, n := range arr {
		if n > 0 {
			contPos++

		} else if n < 0 {
			contNeg++

		} else {
			contZero++

		}

	}

	neg := contNeg / size
	pos := contPos / size
	zero := contZero / size

	fmt.Println(pos)
	fmt.Println(neg)
	fmt.Println(zero)
}
