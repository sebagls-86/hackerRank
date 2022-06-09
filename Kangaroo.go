package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here

	var result string
	if v1 > v2 && ((x2-x1)%(v1-v2)) == 0 || v1 == v2 && x1 == x2 {
		result = "YES"
	} else {
		result = "NO"
	}

	return result
}
