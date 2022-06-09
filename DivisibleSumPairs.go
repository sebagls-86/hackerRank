package main

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var result, j int32

	for i := range ar {
		for j = int32(i) + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				result++
			}
		}
	}

	return result
}
