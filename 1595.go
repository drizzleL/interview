package main

import "math"

func connectTwoGroups(cost [][]int) int {
	minVal := make([]int, len(cost[0]))
	for j := range minVal {
		minVal[j] = cost[0][j]
	}
	for i := 0; i < len(cost); i++ {
		for j := 0; j < len(cost[0]); j++ {
			minVal[j] = min(minVal[j], cost[i][j])
		}
	}
	size2 := len(cost[0])
	size := int(math.Pow(2, float64(size2)))
	cache := make([][]int, len(cost))
	for i := range cache {
		cache[i] = make([]int, size)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var helper func(i int, mask int) int
	helper = func(i int, mask int) (ret int) {
		if i == len(cost) {
			for j := 0; j < len(cost[0]); j++ {
				if mask&(1<<j) == 0 {
					ret += minVal[j]
				}
			}
			return ret
		}
		if cache[i][mask] != -1 {
			return cache[i][mask]
		}
		defer func() {
			cache[i][mask] = ret
		}()
		ret = math.MaxInt32
		for j := 0; j < len(cost[0]); j++ {
			ret = min(ret, cost[i][j]+helper(i+1, mask|1<<j))
		}
		return
	}
	return helper(0, 0)
}
