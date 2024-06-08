package main

func minCut(s string) int {
	dp := make([]int, len(s)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = len(s)
	}
	dp[0] = -1
	palis := make([][]int, len(s))
	for i := range palis {
		palis[i] = make([]int, len(s))
	}

	var isPali func(i, j int) (ret bool)
	isPali = func(i, j int) (ret bool) {
		if palis[i][j] != 0 {
			return palis[i][j] == 1
		}
		defer func() {
			if ret {
				palis[i][j] = 1
			} else {
				palis[i][j] = 2
			}
		}()
		if i >= j {
			return true
		}
		if s[i] != s[j] {
			return false
		}
		return isPali(i+1, j-1)

	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if isPali(i, j) {
				dp[j+1] = min(dp[j+1], dp[i]+1)
			}
		}
	}
	return dp[len(s)]
}
