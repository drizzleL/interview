package main

func minCosts(cost []int) []int {
	ret := make([]int, len(cost))
	ret[0] = cost[0]
	for i := 1; i < len(cost); i++ {
		ret[i] = min(ret[i-1], cost[i])
	}
	return ret
}
