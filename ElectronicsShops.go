package main

func electronicsShop(keyboards []int32, drives []int32, b int32) int32 {
	var result int32 = -1
	var total int32

	for _, k := range keyboards {

		for _, d := range drives {
			total = k + d

			if total <= b && total > result {
				result = total
			}

		}
	}

	return result
}
