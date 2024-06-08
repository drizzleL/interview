package main

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(a+cost[i-2], b+cost[i-1])
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
