package main

func minBitwiseArray(nums []int) []int {
	ret := make([]int, len(nums))
	for i, num := range nums {
		if num&1 == 0 {
			ret[i] = -1
			continue
		}
		left := num & (num + 1)
		ret[i] = left + (num-left)>>1
	}
	return ret
}
