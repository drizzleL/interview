package main

import (
	"math"
)

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	bMin1, bMin2 := make([]int, n+1), make([]int, n+1)
	for i := range bMin1 {
		bMin1[i] = n + 1
		bMin2[i] = n + 1
	}
	for _, p := range conflictingPairs {
		a, b := min(p[0], p[1]), max(p[0], p[1])
		if b < bMin1[a] {
			bMin2[a] = bMin1[a]
			bMin1[a] = b
		} else if b < bMin2[a] {
			bMin2[a] = b
		}
	}
	var ret int
	delCnt := make([]int, n+1)
	for i, right, right2 := n, n, n+1; i >= 1; i-- {
		if bMin1[i] < bMin1[right] {
			right2 = min(right2, bMin1[right])
			right = i
		} else {
			right2 = min(right2, bMin1[i])
		}
		ret += bMin1[right] - i
		delCnt[right] += min(right2, bMin2[right]) - bMin1[right]
	}
	var maxVal int
	for _, v := range delCnt {
		maxVal = max(maxVal, v)
	}
	return int64(ret + maxVal)
}

func maxSubarrays2(n int, conflictingPairs [][]int) int64 {
	bMin1 := make([]int, n+1)
	bMin2 := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bMin1[i] = n + 1
		bMin2[i] = n + 1
	}
	for _, pair := range conflictingPairs {
		a := min(pair[0], pair[1])
		b := max(pair[0], pair[1])
		if bMin1[a] > b {
			bMin2[a] = bMin1[a]
			bMin1[a] = b
		} else if bMin2[a] > b {
			bMin2[a] = b
		}
	}
	res, ib1, b2 := 0, n, math.MaxInt32
	delCount := make([]int, n+1)
	for i := n; i >= 1; i-- {
		if bMin1[ib1] > bMin1[i] {
			b2 = min(b2, bMin1[ib1])
			ib1 = i
		} else {
			b2 = min(b2, bMin1[i])
		}
		res += min(bMin1[ib1], n+1) - i
		delCount[ib1] += min(min(b2, bMin2[ib1]), n+1) - min(bMin1[ib1], n+1)
	}
	maxVal := 0
	for _, v := range delCount {
		maxVal = max(maxVal, v)
	}
	return int64(res + maxVal)
}
