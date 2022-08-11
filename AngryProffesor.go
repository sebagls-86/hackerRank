package main

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var total int32

	for _, alumnos := range a {
		if alumnos <= 0 {
			total++
		}
	}
	if total >= k {
		return "NO"
	}
	return "YES"
}
