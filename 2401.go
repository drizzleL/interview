package main

func longestNiceSubarray(nums []int) int {
	var ret int
	var mask int
	for i, j := 0, 0; j < len(nums); j++ {
		for nums[j]&mask != 0 {
			mask &= ^nums[i]
			i++
		}
		mask |= nums[j]
		ret = max(ret, j-i+1)
	}
	return ret
}
