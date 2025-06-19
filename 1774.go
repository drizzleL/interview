package main

import (
	"math"
)

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	var ret int
	diff := math.MaxInt32
	var helper func(topIdx int, now int)
	helper = func(topIdx int, now int) {
		if now >= target {
			if now-target < diff {
				diff = now - target
				ret = now
			}
			return
		}
		if target-now <= diff {
			diff = target - now
			ret = now
		}
		if topIdx == len(toppingCosts) {
			return
		}
		top := toppingCosts[topIdx]
		helper(topIdx+1, now)
		helper(topIdx+1, now+top)
		helper(topIdx+1, now+top)
	}
	for _, base := range baseCosts {
		helper(0, base)
	}
	return ret
}
