package main

func numberOfGoodSubarraySplits(nums []int) int {
	var ret int
	lastIdx := -1
	for i, num := range nums {
		if num != 1 {
			continue
		}
		if lastIdx == -1 {
			ret = 1
		} else {
			ret *= i - lastIdx
			ret %= 1e9 + 7
		}
		lastIdx = i
	}
	return ret
}
