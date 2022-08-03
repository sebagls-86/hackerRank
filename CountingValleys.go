package main

func countingValleys(steps int32, path string) int32 {
	var counts, currentPos int32
	seaLevel := true

	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			currentPos++
		} else {
			currentPos--
		}

		if seaLevel && currentPos == -1 {
			counts++
		}

		seaLevel = currentPos == 0
	}

	return counts
}
