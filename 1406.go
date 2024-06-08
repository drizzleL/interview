package main

import (
	"math"
)

func stoneGameIII(stoneValue []int) string {
	presum := make([]int, len(stoneValue)+1)
	for i, c := range stoneValue {
		presum[i+1] = presum[i] + c
	}
	cache := make([]int, len(stoneValue))
	for i := range cache {
		cache[i] = math.MinInt32
	}
	var helper func(i int) int
	helper = func(i int) (ret int) {
		if i >= len(stoneValue) {
			return 0
		}
		if cache[i] != math.MinInt32 {
			return cache[i]
		}
		var cnt int
		for k := 1; k <= 3 && i+k-1 < len(stoneValue); k++ {
			cnt += stoneValue[i+k-1]
			cache[i] = max(cache[i], cnt+presum[len(presum)-1]-presum[i+k]-helper(i+k))
		}
		return cache[i]
	}
	aliceTake := helper(0)
	bobTake := presum[len(presum)-1] - aliceTake
	switch {
	case aliceTake == bobTake:
		return "Tie"
	case aliceTake > bobTake:
		return "Alice"
	}
	return "Bob"
}
