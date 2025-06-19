package main

func longestMonotonicSubarray(nums []int) int {
	helper := func(nums []int) int {
		ret := 1
		cnt := 1
		for i := 1; i < len(nums); i++ {
			if nums[i] <= nums[i-1] {
				cnt = 1
				continue
			}
			cnt += 1
			ret = max(ret, cnt)
		}
		return ret
	}
	ret := helper(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return max(ret, helper(nums))
}
