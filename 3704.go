package main

import (
	"strconv"
)

func countNoZeroPairs(n int64) int64 {
	str := strconv.Itoa(int(n))
	b := []byte(str)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	size := len(b)
	dp := make([][2][2][2]int, size+1)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	var helper func(idx int, carry int, za, zb int) int
	helper = func(idx int, carry int, za, zb int) int {
		if idx == size {
			if carry == 0 {
				return 1
			}
			return 0
		}
		if dp[idx][carry][za][zb] != -1 {
			return dp[idx][carry][za][zb]
		}
		d := int(b[idx] - '0')
		var ret int
		for i := 0; i <= 9; i++ {
			if i != 0 && za == 1 {
				continue
			}
			j := (d + 10 - (i + carry)) % 10
			if j != 0 && zb == 1 {
				continue
			}
			if idx == 0 && (i == 0 || j == 0) {
				continue
			}
			nextCarry := (i + j + carry) / 10
			nextZa, nextZb := za, zb
			if i == 0 {
				nextZa = 1
			}
			if j == 0 {
				nextZb = 1
			}
			ret += helper(idx+1, nextCarry, nextZa, nextZb)
		}
		dp[idx][carry][za][zb] = ret
		return ret
	}
	return int64(helper(0, 0, 0, 0))
}
