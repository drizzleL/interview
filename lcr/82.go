package main

import (
	"log"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	log.Println(candidates)
	var ret [][]int
	var helper func(i int, curr []int, tmp int)
	helper = func(i int, curr []int, tmp int) {
		if tmp == target {
			dst := make([]int, len(curr))
			copy(dst, curr)
			ret = append(ret, dst)
			return
		}
		if i == len(candidates) || tmp+candidates[i] > target {
			return
		}
		j := i
		for ; j < len(candidates) && candidates[i] == candidates[j]; j++ {
		}
		helper(j, curr, tmp)
		for m := i; m < j; m++ {
			curr = append(curr, candidates[i])
			tmp += candidates[i]
			helper(j, curr, tmp)
		}
	}
	helper(0, nil, 0)
	return ret
}
