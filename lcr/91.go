package main

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	var a, b, c int
	for i := 0; i < len(costs); i++ {
		a, b, c = costs[i][0]+min(b, c), costs[i][1]+min(a, c), costs[i][2]+min(a, b)
	}
	return min(min(a, b), c)
}
