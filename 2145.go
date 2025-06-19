package main

func numberOfArrays(differences []int, lower int, upper int) int {
	var v, minVal, maxVal int
	for _, diff := range differences {
		v += diff
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}
	gap := maxVal - minVal
	gap2 := upper - lower
	if gap > gap2 {
		return 0
	}
	return gap2 - gap
}
