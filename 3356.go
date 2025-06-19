package main

func minZeroArray(nums []int, queries [][]int) int {
	op := make([]int, len(nums)+1)
	var freq int
	var k int
	for i, num := range nums {
		freq += op[i]
		for num > freq && k < len(queries) {
			q := queries[k]
			op[q[0]] += q[2]
			op[q[1]+1] -= q[2]
			if i >= q[0] && i <= q[1] {
				freq += q[2]
			}
			k += 1
		}
		if num <= freq {
			continue
		}
		return -1
	}
	return k
}
