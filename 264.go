package main

func nthUglyNumber(n int) int {
	nums := []int{1}
	var idx2, idx3, idx5 int
	for i := 1; i < n; i++ {
		num := min(min(nums[idx2]*2, nums[idx3]*3), nums[idx5]*5)
		nums = append(nums, num)
		if num == nums[idx2]*2 {
			idx2 += 1
		}
		if num == nums[idx3]*3 {
			idx3 += 1
		}
		if num == nums[idx5]*5 {
			idx5 += 1
		}
	}
	return nums[len(nums)-1]
}
