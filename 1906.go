package main

func minDifference3(nums []int, queries [][]int) []int {
	dict := make([][101]int, len(nums)+1)
	for i, num := range nums {
		dict[i+1] = dict[i]
		dict[i+1][num] += 1
	}
	ret := make([]int, len(queries))
	for i := range queries {
		q := queries[i]
		a, b := q[0], q[1]
		c := [101]int{}
		ret[i] = -1
		last := -1
		for i := range c {
			c[i] = dict[b+1][i] - dict[a][i]
			if c[i] == 0 {
				continue
			}
			if last != -1 && (ret[i] == -1 || i-last < ret[i]) {
				ret[i] = i - last
			}
			last = i
		}
	}
	return ret
}
