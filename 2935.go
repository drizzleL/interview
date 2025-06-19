package main

import (
	"math/bits"
	"sort"
)

func maximumStrongPairXor(nums []int) int {
	type trieNode struct {
		cnt      [2]int
		children [2]*trieNode
	}
	sort.Ints(nums)
	maxBit := bits.Len(uint(nums[len(nums)-1]))
	root := &trieNode{}
	addVal := func(val int) {
		node := root
		for j := maxBit - 1; j >= 0; j-- {
			bit := (val >> j) & 1
			if node.children[bit] == nil {
				node.children[bit] = &trieNode{}
			}
			node.cnt[bit] += 1
			node = node.children[bit]
		}
	}
	delVal := func(val int) {
		node := root
		for j := maxBit - 1; j >= 0; j-- {
			bit := (val >> j) & 1
			node.cnt[bit] -= 1
			node = node.children[bit]
		}
	}
	var ret int
	check := func(val int) int {
		node := root
		var ret int
		for j := maxBit - 1; j >= 0; j-- {
			bit := ((val >> j) & 1) ^ 1
			if node.cnt[bit] != 0 {
				node = node.children[bit]
				ret |= 1 << j
			} else {
				node = node.children[bit^1]
			}
		}
		return ret
	}
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && nums[j] <= nums[i]*2 {
			addVal(nums[j])
			j += 1
		}
		ret = max(ret, check(nums[i]))
		delVal(nums[i])
	}
	return ret
}
