package main

import (
	"math"
	"sort"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	var pairs [][2]int
	for i, num := range nums {
		pairs = append(pairs, [2]int{i, num})
	}
	sort.Slice(pairs, func(i, j int) bool { // smaller num dealt first
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	h := int(math.Ceil(math.Log2(float64(n)))) + 1
	jumps := make([][]int, n)
	for i := range jumps {
		jumps[i] = make([]int, h)
		for j := range jumps[i] {
			jumps[i][j] = -1
		}
	}
	dict := map[int]int{}
	for i, j := 0, 0; i < len(pairs); i++ {
		p := pairs[i]
		dict[p[0]] = i
		for j+1 < len(pairs) && p[1]+maxDiff >= pairs[j+1][1] {
			j++
		}
		jumps[i][0] = j
	}
	for j := 1; j < h; j++ {
		for i := 0; i < n; i++ {
			jumps[i][j] = jumps[jumps[i][j-1]][j-1]
		}
	}
	var query func(i, j int) int
	query = func(i, j int) int {
		if i == j {
			return 0
		}
		if i > j {
			return query(j, i)
		}
		if jumps[i][0] >= j {
			return 1
		}
		var m int
		for m < h && jumps[i][m] < j {
			m += 1
		}
		if m == h {
			return -1
		}
		m -= 1
		return 1<<m + query(jumps[i][m], j)

	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		ret[i] = query(dict[q[0]], dict[q[1]])
	}
	return ret
}
