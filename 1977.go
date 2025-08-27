package main

func numberOfCombinations(num string) int {
	n := len(num)
	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}
	for j := n - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; i-- {
			if num[i] == num[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			}
		}
	}
	cmp := func(i1, i2 int, threshold int) bool {
		if lcs[i1][i2] >= threshold {
			return true
		}
		v := lcs[i1][i2]
		return num[i1+v] <= num[i2+v]
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	mod := int(1e9 + 7)
	for j := 0; j < n; j++ {
		for size := 1; size <= j+1; size++ {
			dp[j][size] += dp[j][size-1]
			i := j - size + 1
			if num[i] == '0' {
				continue
			}
			var cnt int
			if i == 0 {
				cnt = 1
			} else {
				preI, preJ := i-size, i-1
				if preI >= 0 && num[preI] != '0' && cmp(preI, i, size) {
					cnt += (dp[preJ][size] - dp[preJ][size-1] + mod) % mod
				}
				prevLenLimit := size - 1
				if prevLenLimit > i {
					prevLenLimit = i
				}
				cnt += dp[preJ][prevLenLimit]
			}
			dp[j][size] += cnt
			dp[j][size] %= mod
		}
	}
	return dp[n-1][n]
}
