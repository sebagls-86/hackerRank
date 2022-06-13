package main

import "strconv"

func dayOfProgrammer(year int32) string {
	// Write your code here
	finalDay := ""

	s := strconv.FormatInt(int64(year), 10)

	if year == 1918 {
		finalDay = "26.09." + s

	} else if year < 1918 && year%4 == 0 || year > 1918 && year%400 == 0 || year%4 == 0 && year%100 != 0 {
		finalDay = "12.09." + s
	} else {
		finalDay = "13.09." + s
	}

	return finalDay
}
