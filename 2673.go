package main

func minIncrements(n int, cost []int) int {
	maxChild := make([]int, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i*2+1 >= n {
			return 0
		}
		l, r := dfs(i*2+1), dfs(i*2+2)
		maxChild[i] = max(l, r)
		return cost[i] + maxChild[i]
	}
	var ret int
	var helper func(i int, v int)
	helper = func(i int, v int) {
		if i*2+1 >= n {
			return
		}
		v -= cost[i]
		v -= maxChild[i]
		ret += v
		helper(i*2+1, maxChild[i])
		helper(i*2+2, maxChild[i])
	}
	maxVal := dfs(0)
	helper(0, maxVal)
	return ret
}
