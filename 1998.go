package main

import "sort"

func gcdSort(nums []int) bool {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	getPrimes := func(x int) []int {
		var ret []int
		for v := 2; x != 1 && v*v <= x; v++ {
			if x%v != 0 {
				continue
			}
			ret = append(ret, v)
			for x%v == 0 {
				x /= v
			}
		}
		if x != 1 {
			ret = append(ret, x)
		}
		return ret
	}
	parent := make([]int, len(nums))
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	groupDict := map[int]int{}
	for i, num := range nums {
		for _, p := range getPrimes(num) {
			if oldGroup, ok := groupDict[p]; ok {
				union(i, oldGroup)
			} else {
				groupDict[p] = i
			}
		}
	}
	groupVals := map[int][]int{}
	for i, num := range nums {
		p := find(i)
		groupVals[p] = append(groupVals[p], num)
	}
	for _, g := range groupVals {
		sort.Ints(g)
	}
	vals := make([]int, len(nums))
	copy(vals, nums)
	sort.Ints(vals)
	for i := range nums {
		p := find(i)
		val := groupVals[p][0]
		if vals[i] != val {
			return false
		}
		groupVals[p] = groupVals[p][1:]
	}
	return true
}
