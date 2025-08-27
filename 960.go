package main

func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	ret := n
	dp := make([]int, n+1)
	check := func(j, j2 int) bool {
		for i := 0; i < m; i++ {
			if strs[i][j] < strs[i][j2] {
				return false
			}
		}
		return true
	}
	dp[0] = 1
	for j := 1; j < n; j++ {
		for j2 := 0; j2 < j; j2++ {
			if check(j, j2) {
				dp[j] = max(dp[j], dp[j2]+1)
			}
		}
		ret = min(ret, n-dp[j])
	}
	return ret
}
