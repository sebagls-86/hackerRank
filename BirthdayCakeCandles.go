package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var result, max int32

	for _, value := range candles {

		if value > max {
			max = value
			result = 1
		} else if value == max {
			result++
		}
	}

	return result
}
