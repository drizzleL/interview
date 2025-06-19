package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	nums := []int{1}
	idxs := make([]int, len(primes))
	for i := 1; i < n; i++ {
		num := math.MaxInt32
		for i, idx := range idxs {
			num = min(num, nums[idx]*primes[i])
		}
		nums = append(nums, num)
		for i, idx := range idxs {
			if num == nums[idx]*primes[i] {
				idxs[i] += 1
			}
		}
	}
	return nums[len(nums)-1]
}
