package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	var helper func(now []int, sum int, i int)
	helper = func(now []int, sum int, i int) {
		if sum == target {
			cpy := make([]int, len(now))
			copy(cpy, now)
			ret = append(ret, cpy)
			return
		}
		if sum > target || i == len(candidates) {
			return
		}
		v := candidates[i]
		j := i
		for ; j < len(candidates) && candidates[j] == v; j++ {
		}
		helper(now, sum, j)
		for ; i < j; i++ {
			now = append(now, v)
			sum += v
			helper(now, sum, j)
		}
	}
	helper(nil, 0, 0)
	return ret
}
