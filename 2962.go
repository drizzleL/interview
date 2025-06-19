package main

func countSubarrays2(nums []int, k int) int64 {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	var cnt int
	var ret int
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && cnt < k {
			if nums[j] == maxVal {
				cnt += 1
			}
			j++
		}
		if cnt < k {
			break
		}
		ret += len(nums) - j + 1
		if nums[i] == maxVal {
			cnt -= 1
		}
	}
	return int64(ret)

}
