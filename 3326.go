package main

func minOperations10(nums []int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]int, maxVal+1)
	for i := 2; i < len(sieve); i++ {
		if sieve[i] > 0 { // marked before
			continue
		}
		for j := i; j < len(sieve); j += i {
			if sieve[j] == 0 {
				sieve[j] = i
			}
		}
	}
	var ret int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			continue
		}
		nums[i] = sieve[nums[i]]
		ret += 1
		if nums[i] > nums[i+1] {
			return -1
		}
	}
	return ret
}
