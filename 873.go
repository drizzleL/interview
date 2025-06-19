package main

func lenLongestFibSubseq(arr []int) int {
	dict := map[int]int{}
	for i, v := range arr {
		dict[v] = i
	}
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	var ret int
	for j := len(arr) - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; i-- {
			k, ok := dict[arr[i]+arr[j]]
			dp[i][j] = 2
			if ok {
				dp[i][j] = dp[j][k] + 1
			}
			ret = max(ret, dp[i][j])
		}
	}
	if ret == 2 {
		return 0
	}
	return ret
}
