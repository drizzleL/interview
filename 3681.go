package main

import "math/bits"

func maxXorSubsequences(nums []int) int {
	type trieNode struct {
		children [2]*trieNode
	}
	var maxSize int
	for _, num := range nums {
		maxSize = max(maxSize, bits.Len(uint(num)))
	}
	basis := make([]int, maxSize)
	for _, num := range nums {
		for i := maxSize - 1; i >= 0; i-- {
			if num&(1<<i) == 0 {
				continue
			}
			if basis[i] == 0 {
				basis[i] = num
				break
			}
			num ^= basis[i]
		}
	}
	var ret int
	for i := maxSize - 1; i >= 0; i-- {
		ret = max(ret, ret^basis[i])
	}
	return ret
}
