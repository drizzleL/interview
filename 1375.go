package main

func numTimesAllBlue(flips []int) int {
	var maxIdx int
	var ret int
	for i, f := range flips {
		maxIdx = max(maxIdx, f)
		if i+1 == maxIdx {
			ret += 1
		}
	}
	return ret
}
