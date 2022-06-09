package main

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var i, j, result int32

	for i = 1; i <= 100; i++ {
		factor := true

		for j = 0; j < int32(len(a)); j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}
		if factor {
			for j = 0; j < int32(len(b)); j++ {
				if b[j]%i != 0 {
					factor = false
					break
				}
			}
		}

		if factor {
			result++
		}
	}

	return result
}
