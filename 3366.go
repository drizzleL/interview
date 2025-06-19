package main

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	var helper func(i int, op1, op2 int) int
	cache := make([][][]int, len(nums)+1)
	for i := range cache {
		cache[i] = make([][]int, op1+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, op2+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	helper = func(i int, op1, op2 int) (ret int) {
		if c := cache[i][op1][op2]; c != -1 {
			return c
		}
		defer func() {
			cache[i][op1][op2] = ret
		}()
		if i == len(nums) {
			return 0
		}
		ret = nums[i] + helper(i+1, op1, op2)
		if op1 > 0 && op2 > 0 {
			if nums[i] >= k*2-1 {
				ret = min(ret, (nums[i]+1)/2-k+helper(i+1, op1-1, op2-1))
			}
			if nums[i] >= k {
				ret = min(ret, (nums[i]-k+1)/2+helper(i+1, op1-1, op2-1))
			}
		}
		if op1 > 0 {
			ret = min(ret, (nums[i]+1)/2+helper(i+1, op1-1, op2))
		}
		if op2 > 0 && nums[i] >= k {
			ret = min(ret, nums[i]-k+helper(i+1, op1, op2-1))
		}
		return
	}
	return helper(0, op1, op2)
}
