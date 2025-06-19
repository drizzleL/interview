package main

func isZeroArray(nums []int, queries [][]int) bool {
	op := make([]int, len(nums)+1)
	for _, q := range queries {
		op[q[0]] += 1
		op[q[1]+1] -= 1
	}
	var freq int
	for i, num := range nums {
		freq += op[i]
		if freq < num {
			return false
		}
	}
	return true
}
