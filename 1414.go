package main

func findMinFibonacciNumbers(k int) int {
	nums := []int{1, 1}
	for nums[len(nums)-1] < k {
		nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
	}
	var ret int
	for k != 0 {
		for nums[len(nums)-1] > k {
			nums = nums[:len(nums)-1]
		}
		k -= nums[len(nums)-1]
		ret += 1
	}
	return ret
}
