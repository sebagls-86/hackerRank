package main

func theHardleRacek(k int32, height []int32) int32 {

	var result int32

	bigger := height[0]
	for _, number := range height {
		if number > bigger {
			bigger = number
		}

	}

	result = bigger - k

	if result < 0 {
		result = 0
	}

	return result
}
