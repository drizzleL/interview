package main

func longestSubarray2(nums []int) int {
	var maxVal, cnt, ret int
	for _, num := range nums {
		if num < maxVal {
			cnt = 0
			continue
		}
		if num == maxVal {
			cnt += 1
			ret = max(ret, cnt)
			continue
		}
		maxVal = num
		cnt = 1
		ret = max(ret, cnt)
	}
	return ret
}
