package main

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var sum int64

	for i := 0; i < len(ar); i++ {
		sum = sum + ar[i]
	}

	return sum

}
