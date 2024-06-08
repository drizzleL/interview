package main

import (
	"math"
)

func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	cache := map[[2]int]int{}
	var helper func(i int, d int) int
	helper = func(i int, d int) (ret int) {
		if v, ok := cache[[2]int{i, d}]; ok {
			return v
		}
		defer func() {
			cache[[2]int{i, d}] = ret
		}()
		if d == 1 {
			for ; i < len(jobDifficulty); i++ {
				ret = max(ret, jobDifficulty[i])
			}
			return
		}
		ret = math.MaxInt32
		var val int
		for j := i; j <= len(jobDifficulty)-d; j++ {
			val = max(val, jobDifficulty[j])
			ret = min(ret, val+helper(j+1, d-1))
		}
		return ret
	}
	return helper(0, d)
}
