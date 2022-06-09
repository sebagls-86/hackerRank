package main

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var result []int32
	var max, min, totalMax, totalMin int32

	max = scores[0]
	min = scores[0]

	for _, score := range scores {
		if score > max {
			max = score
			totalMax++
		}
	}

	for _, score := range scores {

		if score < min {
			min = score
			totalMin++
		}
	}

	result = append(result, totalMax, totalMin)

	return result
}
