package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	var miniSum, maxSum int64

	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
	var sliceMin []int32 = arr[0:4]

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var sliceMax []int32 = arr[1:5]

	for _, value := range sliceMin {
		miniSum += int64(value)
	}

	for _, value := range sliceMax {
		maxSum += int64(value)
	}

	fmt.Print(miniSum, maxSum)

}
