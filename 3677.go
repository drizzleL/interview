package main

import (
	"math/bits"
	"strconv"
)

func countBinaryPalindromes(n int64) int {
	if n <= 1 {
		return int(n) + 1
	}
	size := bits.Len(uint(n))
	dp := make([]int, size)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < size; i++ {
		dp[i] = 2 * dp[i-2]
	}
	arr := make([]int, size)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < size; i++ {
		arr[i] = dp[i-2]
	}
	var ret int
	for i := 1; i < size; i++ {
		ret += arr[i]
	}
	s := strconv.FormatInt(n, 2)
	ret += 1
	for i, j := len(s)/2-1, (len(s)+1)/2; i >= 0; i, j = i-1, j+1 {
		if s[i] == s[j] {
			continue
		}
		if s[i] > s[j] {
			ret -= 1
		}
		break
	}
	for i, j := 1, len(s)-2; i <= j; i, j = i+1, j-1 {
		if s[i] == '0' { // make it 0, add up
			continue
		}
		leftSize := max(j-i-1, 0)
		ret += dp[leftSize]
	}
	return ret
}
