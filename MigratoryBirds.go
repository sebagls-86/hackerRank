package main

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var count, parcial, id int
	var resultado int32

	ar := [5]int32{1, 2, 3, 4, 5}

	for i := range ar {
		count = 0
		for j := range arr {
			if ar[i] == arr[j] {
				count++
			}
		}
		if count > parcial {
			parcial = count
			resultado = ar[i]
		}
		if count == parcial && i < id {
			resultado = ar[i]
		}

	}

	return int32(resultado)

}
