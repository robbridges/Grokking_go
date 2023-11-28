package main



func LargestAltitude(gain []int) int {
	altitude := 0
	maxAltitude := 0
	for _, num := range gain {
		altitude += num
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}
	return maxAltitude
}