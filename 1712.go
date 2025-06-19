package main

import (
	"sort"
)

func waysToSplit(nums []int) int {
	var sum int
	sums := make([]int, len(nums))
	for i, num := range nums {
		sum += num
		sums[i] = num
		if i != 0 {
			sums[i] += sums[i-1]
		}
	}
	var ret int
	for i := 0; i < len(nums)-2; i++ {
		sum1 := sums[i]
		idx := sort.SearchInts(sums, sum1*2)
		if idx >= len(sums)-1 {
			break
		}
		idx = max(idx, i+1)
		idx2 := sort.SearchInts(sums, sum1+(sum-sum1)/2+1)
		idx2 = min(len(sums)-1, idx2)
		if idx2 < idx {
			break
		}
		ret += idx2 - idx
		ret %= 1e9 + 7
	}
	return ret
}
