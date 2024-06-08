package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	var ret int
	for _, str := range strs {
		var c1, c2 int
		for _, c := range str {
			switch c {
			case '0':
				c1++
			case '1':
				c2++
			}
		}
		for i := m; i >= c1; i-- {
			for j := n; j >= c2; j-- {
				if dp[i-c1][j-c2] == -1 {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[i-c1][j-c2]+1)
				ret = max(ret, dp[i][j])
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
