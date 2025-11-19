package main

import "sort"

func maxAlternatingSum(nums []int, swaps [][]int) int64 {
	parents := make([]int, len(nums))
	for i := range parents {
		parents[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parents[pb] = pa
	}
	for _, swap := range swaps {
		union(swap[0], swap[1])
	}
	group := map[int][]int{}
	for i := range parents {
		group[find(i)] = append(group[find(i)], i)
	}
	var ret int
	for _, mems := range group {
		var vals []int
		var odd int
		for _, i := range mems {
			vals = append(vals, nums[i])
			odd += i % 2
		}
		sort.Ints(vals)
		for i := 0; i < odd; i++ {
			ret -= vals[i]
		}
		for i := odd; i < len(vals); i++ {
			ret += vals[i]
		}
	}
	return int64(ret)
}
