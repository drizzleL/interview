package main

import (
	"math"
)

func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(edges) + 1
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	seen := make([]bool, n)
	dirs := make([][]int, n)
	nodes := []int{0}
	for len(nodes) != 0 {
		var nextNodes []int
		for _, node := range nodes {
			seen[node] = true
			for _, next := range dict[node] {
				if seen[next] {
					continue
				}
				dirs[node] = append(dirs[node], next)
				nextNodes = append(nextNodes, next)
			}
		}
		nodes = nextNodes
	}
	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, k)
		for j := range cache[i] {
			cache[i][j] = make([]int, 2)
			for m := range cache[i][j] {
				cache[i][j][m] = math.MinInt32
			}
		}
	}
	var dfs func(i int, k int, flip int) int
	dfs = func(i int, now int, flip int) (ret int) {
		if i == n {
			return 0
		}
		if cache[i][now][flip] != math.MinInt32 {
			return cache[i][now][flip]
		}
		defer func() {
			cache[i][now][flip] = ret
		}()
		ret = math.MinInt32
		factor := 1
		if flip == 1 {
			factor *= -1
		}
		val := factor * nums[i]
		for _, next := range dirs[i] {
			val += dfs(next, max(0, now-1), flip)
		}
		ret = max(ret, val)
		if now == 0 {
			val = -factor * nums[i]
			for _, next := range dirs[i] {
				val += dfs(next, k-1, 1-flip)
			}
			ret = max(ret, val)
		}
		return
	}
	return int64(dfs(0, 0, 0))
}
