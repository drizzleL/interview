package main

import (
	"math"
)

func minimumCost2(s string) int64 {
	prefixSum := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		prefixSum[i] = prefixSum[i-1]
		if s[i] != s[i-1] {
			prefixSum[i] += i
		}
	}
	suffixSum := make([]int, len(s))
	for i := len(s) - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1]
		if s[i] != s[i+1] {
			suffixSum[i] += len(s) - i - 1
		}
	}
	ret := math.MaxInt64
	for i := 0; i < len(s); i++ {
		ret = min(ret, prefixSum[i]+suffixSum[i])
	}
	return int64(ret)
}
