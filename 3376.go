package main

import "math"

func findMinimumTime(strength []int, k int) int {
	var helper func(x int, flag int) int
	cache := map[[2]int]int{}
	helper = func(x int, flag int) (ret int) {
		if c, ok := cache[[2]int{x, flag}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{x, flag}] = ret
		}()
		if flag == (1<<len(strength))-1 {
			return 0
		}
		ret = math.MaxInt32
		for i := 0; i < len(strength); i++ {
			if flag&(1<<i) != 0 {
				continue
			}
			ret = min(ret, (strength[i]+x-1)/x+helper(x+k, flag|1<<i))
		}
		return
	}
	return helper(1, 0)
}
