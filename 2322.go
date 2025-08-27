package main

import (
	"math"
	"sort"
)

func minimumScore(nums []int, edges [][]int) int {
	var xorSum int
	for _, num := range nums {
		xorSum ^= num
	}
	dict := make([][]int, len(nums))
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	xors := make([]int, len(nums))
	children := make([][]bool, len(nums))
	for i := range children {
		children[i] = make([]bool, len(nums))
	}
	var dfs func(i int, pre int) (int, []int)
	dfs = func(i int, pre int) (int, []int) {
		xors[i] = nums[i]
		childIds := []int{i}
		for _, child := range dict[i] {
			if child == pre {
				continue
			}
			childXor, childs := dfs(child, i)
			xors[i] ^= childXor
			for _, childId := range childs {
				children[childId][i] = true
				childIds = append(childIds, childId)
			}
		}
		return xors[i], childIds
	}
	dfs(0, -1)
	helper := func(vals []int) int {
		sort.Ints(vals)
		return vals[len(vals)-1] - vals[0]
	}
	ret := math.MaxInt32
	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			a, b := edges[i][0], edges[i][1]
			if children[b][a] {
				a, b = b, a
			}
			c, d := edges[j][0], edges[j][1]
			if children[d][c] {
				c, d = d, c
			}
			v1, v2 := xors[a], xors[c]
			var vals []int
			if children[a][c] {
				vals = []int{v1, v2 ^ v1, v2 ^ xorSum}
			} else if children[c][a] {
				vals = []int{v2, v2 ^ v1, v1 ^ xorSum}
			} else {
				vals = []int{v1, v2, v1 ^ v2 ^ xorSum}
			}
			ret = min(ret, helper(vals))
		}
	}
	return ret
}
