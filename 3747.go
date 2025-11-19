package main

import "strconv"

func countDistinct2(n int64) int64 {
	str := strconv.Itoa(int(n))
	dp := make([]int, len(str))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		dp[i] = 9 * dp[i-1]
	}
	var ret int
	for i := 1; i < len(dp); i++ {
		ret += dp[i]
	}
	for i := 0; i < len(dp); i++ {
		if str[i] == '0' {
			continue
		}
		ret += int(str[i]-'1') * dp[len(str)-i-1]
	}
	for _, c := range str {
		if c == '0' {
			return int64(ret)
		}
	}
	return int64(ret) + 1
}
