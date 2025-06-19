package main

func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	cache := map[[3]int]int{}
	var helper func(i, j, k int) int
	helper = func(i, j, k int) (ret int) {
		if j-i+1 < 2 {
			return 0
		}
		if c, ok := cache[[3]int{i, j, k}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{i, j, k}] = ret
		}()
		if nums[i]+nums[i+1] == k {
			ret = max(ret, helper(i+2, j, k))
		}
		if nums[i]+nums[j] == k {
			ret = max(ret, helper(i+1, j-1, k))
		}
		if nums[j]+nums[j-1] == k {
			ret = max(ret, helper(i, j-2, k))
		}
		return
	}
	ret := max(helper(0, len(nums)-1, nums[0]+nums[1]), helper(0, len(nums)-1, nums[0]+nums[len(nums)-1]))
	ret = max(ret, helper(0, len(nums)-1, nums[len(nums)-1]+nums[len(nums)-2]))
	return ret
}
