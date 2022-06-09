package main

import "math"

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here

	d1 := []int32{}
	d2 := []int32{}
	var d1Sum, d2Sum int32

	for x := range arr {
		for y := range arr {
			if x == y {
				d1 = append(d1, arr[x][y])
				d1Sum += arr[x][y]
			}

			if x == len(arr)-y-1 {
				d2 = append(d2, arr[x][y])
				d2Sum += arr[x][y]
			}
		}
	}

	return int32(math.Abs(float64(d1Sum) - float64(d2Sum)))
}
