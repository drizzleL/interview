package main

import (
	"math/big"
)

func countNumbers(l string, r string, b int) int {
	toNum := func(s string) string {
		v := big.Int{}
		v2, ok := v.SetString(s, 10)
		if !ok {
			return ""
		}
		return v2.Text(b)
	}
	l2, r2 := toNum(l), toNum(r)
	dp := make([][]int, len(r2)+1)
	for i := range dp {
		dp[i] = make([]int, b)
	}
	for j := b - 1; j >= 0; j-- {
		dp[0][j] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < b; j++ {
			for k := j; k < b; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}
	helper := func(s string) int {
		var ret, pre int
		for i := 0; i < len(s); i++ {
			if i != 0 && s[i] < s[i-1] {
				break
			}
			d := int(s[i] - '0')
			if i == len(s)-1 {
				ret += d - pre + 1
				break
			}
			leftSize := len(s) - i - 1
			for j := pre; j < d; j++ {
				ret += dp[leftSize][j]
				ret %= 1e9 + 7
			}
			pre = max(pre, d)
		}
		return ret % (1e9 + 7)
	}
	ret := helper(r2) - helper(l2)
	if ret < 0 {
		ret += 1e9 + 7
	}
	check := func(s string) bool {
		for i := 1; i < len(s); i++ {
			if s[i] < s[i-1] {
				return false
			}
		}
		return true
	}
	if check(l2) {
		ret += 1
	}
	return ret % (1e9 + 7)
}
