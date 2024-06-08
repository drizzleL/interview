package main

func countSubarrays2(nums []int, k int) int64 {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	var cnt int
	var l int
	var ret int
	for r, num := range nums {
		if num == maxVal {
			cnt += 1
		}
		for cnt == k {
			ret += len(nums) - r
			if nums[l] == maxVal {
				cnt -= 1
			}
			l++
		}
	}
	return int64(ret)
}
