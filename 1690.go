package main

func stoneGameVII(stones []int) int {
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	sums := make([]int, len(stones)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + stones[i-1]
	}
	var helper func(i, j int) int
	helper = func(i, j int) (ret int) {
		if i >= j {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		defer func() {
			dp[i][j] = ret
		}()
		rangeSum := sums[j+1] - sums[i]
		a, b := rangeSum-stones[i], rangeSum-stones[j]
		ret = max(a-helper(i+1, j), b-helper(i, j-1))
		return
	}
	return helper(0, len(stones)-1)
}
