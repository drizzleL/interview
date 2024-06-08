package main

func numberOfArrays(differences []int, lower int, upper int) int {
	var minVal, maxVal, val int
	for _, diff := range differences {
		val += diff
		minVal = min(minVal, val)
		maxVal = max(maxVal, val)
	}
	gap := maxVal - minVal
	if gap > upper-lower {
		return 0
	}
	return upper - lower - gap + 1
}
