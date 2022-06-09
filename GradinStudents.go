package main

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var result []int32

	for _, grade := range grades {

		if grade < 38 {
			result = append(result, grade)
		} else {
			if grade%5 > 2 {
				grade = ((grade / 5) + 1) * 5
				result = append(result, grade)
			} else {
				result = append(result, grade)
			}
		}
	}

	return result

}
