package main

func isCovered(ranges [][]int, left int, right int) bool {
	seen := make([]bool, right-left+1)
	for _, r := range ranges {
		if r[0] > right {
			continue
		}
		if r[1] < left {
			continue
		}
		a, b := max(r[0], left), min(r[1], right)
		for v := a; v <= b; v++ {
			seen[v-left] = true
		}
	}
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
