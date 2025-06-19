package main

import (
	"math"
	"math/bits"
)

func minimumDifference(nums []int, k int) int {
	var bitSize int
	for _, num := range nums {
		bitSize = max(bitSize, bits.Len(uint(num)))
	}
	ret := math.MaxInt32
	var val int
	dict := make([]int, bitSize)
	for l, r := 0, 0; r < len(nums); {
		for r < len(nums) && val < k {
			if l != r {
				ret = min(ret, k-val)
			}
			for bit := 0; bit < bitSize; bit++ {
				if nums[r]&(1<<bit) != 0 {
					dict[bit] += 1
				}
			}
			val |= nums[r]
			r++
		}
		if val < k {
			ret = min(ret, k-val)
			break
		}
		if val == k { // early return
			return 0
		}
		for val > k && l < r {
			ret = min(ret, val-k)
			for bit := 0; bit < bitSize; bit++ {
				if nums[l]&(1<<bit) != 0 {
					dict[bit] -= 1
					if dict[bit] == 0 {
						val ^= 1 << bit
					}
				}
			}
			l++
		}
		if l != r {
			ret = min(ret, abs(k-val))
		}
	}
	return ret
}
