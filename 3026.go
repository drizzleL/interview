package main

import "math"

func maximumSubarraySum2(nums []int, k int) int64 {
	dict := map[int]int{}
	var presum int
	ret := math.MinInt32
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			dict[num] = min(dict[num], presum)
		} else {
			dict[num] = presum
		}
		presum += num
		if sum, ok := dict[num+k]; ok {
			ret = max(ret, presum-sum)
		}
		if sum, ok := dict[num-k]; ok {
			ret = max(ret, presum-sum)
		}
	}
	if ret == math.MinInt32 {
		return 0
	}
	return int64(ret)
}
