package main

func longestSubarray4(nums []int) int {
	var ret, last, curr int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			last = curr
			curr = 0
		} else {
			curr++
		}
		ret = max(ret, last+curr)
	}
	if curr == len(nums) {
		ret -= 1
	}
	return ret
}
