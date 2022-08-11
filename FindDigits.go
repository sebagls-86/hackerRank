package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findDigits(n int32) int32 {
	// Write your code here
	var count int32

	count = 0

	//str1 := strconv.Itoa(n)

	str1 := strconv.FormatInt(int64(n), 10)

	fmt.Println(str1)

	//str1 := "12"

	res1 := strings.Split(str1, "")

	fmt.Println("\nResult ", res1)

	for i := 0; i < len(res1); i++ {

		intVar, _ := strconv.Atoi(res1[i])

		fmt.Println("sub I ", intVar)

		if int32(intVar) != 0 && n%int32(intVar) == 0 {
			count += 1
		}

	}

	fmt.Println(count)
	return count
}
