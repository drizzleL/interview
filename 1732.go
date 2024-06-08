package main

func largestAltitude(gain []int) int {
	var ret int
	var pre int
	for _, g := range gain {
		pre += g
		ret = max(ret, pre)
	}
	return ret
}
