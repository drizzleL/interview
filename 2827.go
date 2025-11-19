package main

import (
	"math"
	"strconv"
)

func numberOfBeautifulIntegers(low int, high int, k int) int {
	size := len(strconv.Itoa(high))
	offset := size
	dp := make([][][]int, size)
	for i := range dp {
		dp[i] = make([][]int, k)
		for j := range dp[i] {
			dp[i][j] = make([]int, offset*2+1)
		}
	}
	mod := int(1e9 + 7)
	dp[0][0][offset] = 1 // size, mod, odd - even
	for i := 1; i < size; i++ {
		for prevRem := 0; prevRem < k; prevRem++ {
			for preDiff := -(i - 1); preDiff <= i-1; preDiff++ {
				if dp[i-1][prevRem][offset+preDiff] == 0 {
					continue
				}
				for j := 0; j <= 9; j++ {
					newRem := (prevRem*10 + j) % k
					newDiff := preDiff + (j%2)*2 - 1
					dp[i][newRem][offset+newDiff] += dp[i-1][prevRem][offset+preDiff]
				}
			}
		}
	}
	arr := make([]int, size)
	for i := 1; i < size; i++ {
		arr[i] = (dp[i][0][offset] - dp[i-1][0][offset+1] + mod) % mod
	}
	helper := func(x int) int {
		var ret, pre, delta int
		s := strconv.Itoa(x)
		size := len(s)
		for i := 1; i < size; i++ {
			ret += arr[i]
		}
		base := int(math.Pow(10, float64(len(s)-1)))
		for i := 0; i < len(s); i++ {
			d := int(s[i] - '0')
			for j := 0; j < d; j++ {
				if i == 0 && j == 0 {
					continue
				}
				delta2 := delta + (j%2)*2 - 1
				oldMod := (k - (pre+j*base)%k) % k
				ret += dp[len(s)-i-1][oldMod][offset-delta2]
				ret %= 1e9 + 7
			}
			pre += d * base
			delta += (d%2)*2 - 1
			base /= 10
		}
		if x%k == 0 {
			var diff int
			for i := 0; i < len(s); i++ {
				diff += int(s[i]-'0')%2*2 - 1
			}
			if diff == 0 {
				ret += 1
			}
		}
		return ret
	}
	return helper(high) - helper(low-1)
}
