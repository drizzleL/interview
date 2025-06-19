package main

func countArrays(original []int, bounds [][]int) int {
	if len(bounds) == 0 {
		return 0
	}
	b := bounds[0]
	for i := 1; i < len(original); i++ {
		gap := original[i] - original[i-1]
		b[0] += gap
		b[1] += gap
		b[0] = max(b[0], bounds[i][0])
		b[1] = min(b[1], bounds[i][1])
		if b[0] > b[1] {
			return 0
		}
	}
	return b[1] - b[0] + 1
}
