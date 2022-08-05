package main

import "math"

func catAndMouse(x int32, y int32, z int32) string {

	catA := "Cat A"
	catB := "Cat B"
	mouse := "Mouse C"

	distCatA := math.Abs(float64(x) - float64(z))
	distCatB := math.Abs(float64(y) - float64(z))

	if distCatA < distCatB {
		return catA
	} else if distCatB < distCatA {
		return catB
	} else {
		return mouse
	}

}
