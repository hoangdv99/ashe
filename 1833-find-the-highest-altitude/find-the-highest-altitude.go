func largestAltitude(gain []int) int {
	max := 0
	currentAltitude := 0
	for _, v := range gain {
		currentAltitude += v
		if max < currentAltitude {
			max = currentAltitude
		}
	}
	return max
}
