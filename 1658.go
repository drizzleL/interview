package main

import (
	"math"
)

func minOperations2(nums []int, x int) int {
	dict := map[int]int{}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > x {
			break
		}
		dict[sum] = i + 1
	}
	if sum < x {
		return -1
	}
	sum = 0
	ret := math.MaxInt32
	if dict[x] != 0 {
		ret = dict[x]
	}
	for i, j := 1, len(nums)-1; j >= 0; i, j = i+1, j-1 {
		sum += nums[j]
		if sum > x {
			break
		}
		if sum != x && dict[x-sum] == 0 {
			continue
		}
		ret = min(ret, dict[x-sum]+i)
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
