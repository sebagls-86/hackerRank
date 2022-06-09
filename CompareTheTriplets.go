package main

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var ar, br int32

	var result []int32
	for i := 0; i < 3; i++ {
		//compare here
		if a[i] > b[i] {
			ar++
		} else if b[i] > a[i] {
			br++
		}
	}

	result = append(result, ar, br)

	return result
}
