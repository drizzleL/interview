package main

import (
	"math"
)

func minSizeSubarray(nums []int, target int) int {
	presum := map[int]int{0: 0}
	var sum int
	ret := math.MaxInt32
	for i, num := range nums {
		sum += num
		if cnt, ok := presum[sum-target]; ok {
			ret = min(ret, i+1-cnt)
		}
		presum[sum] = i + 1
	}
	totalCnt := int(math.Ceil(float64(target) / float64(sum)))
	var midCnt int
	switch totalCnt {
	case 1, 2:
	case 3:
		midCnt = 1
	default:
		midCnt = totalCnt - 2
	}
	tmp := midCnt * len(nums)
	target -= midCnt * sum

	var tailsum int
	for j := len(nums) - 1; j >= 0; j-- {
		tailsum += nums[j]
		if cnt, ok := presum[target-tailsum]; ok {
			ret = min(ret, cnt+len(nums)-j)
		}
		if cnt, ok := presum[target-sum-tailsum]; ok {
			ret = min(ret, cnt+len(nums)-j+len(nums))
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret + tmp
}
