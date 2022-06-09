package main

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var sum, sp, i, j int32

	for i = 0; i < int32(len(s)); i++ {
		sp = 0

		if i+m > int32(len(s)) {
			return sum
		}

		for j = 0; j < m; j++ {
			sp = sp + s[i+j]
		}

		if sp == d {
			sum++
		}
	}

	return sum

}
