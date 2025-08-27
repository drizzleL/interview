package main

import "math"

func maximizeXorAndXor(nums []int) int64 {
	var xorSum, maxVal int
	for _, num := range nums {
		xorSum ^= num
		maxVal = max(maxVal, num)
	}
	var ret int
	bitSize := int(math.Log2(float64(maxVal))) + 1
	for mask := 0; mask < 1<<len(nums); mask++ {
		var xor int
		and := 1<<bitSize - 1
		var nums2 []int
		for i := 0; i < len(nums); i++ {
			if mask&(1<<i) == 0 {
				nums2 = append(nums2, nums[i])
				continue
			}
			xor ^= nums[i]
			and &= nums[i]
		}
		if mask == 0 {
			and = 0
		}
		u := xorSum ^ xor
		basis := make([]int, bitSize)
		for _, x := range nums2 {
			x &= ^u
			for j := bitSize - 1; j >= 0; j-- {
				if x&(1<<j) == 0 {
					continue
				}
				if basis[j] == 0 {
					basis[j] = x
					break
				}
				x ^= basis[j]
			}
		}
		var maxXor int
		for j := bitSize - 1; j >= 0; j-- {
			maxXor = max(maxXor, maxXor^basis[j])
		}
		ret = max(ret, and+u+2*maxXor)
	}
	return int64(ret)
}
