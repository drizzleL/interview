package main

func splitArray(nums []int) int64 {
	primes := make([]bool, len(nums)+1)
	for i := 2; i <= len(nums); i++ {
		if primes[i] {
			continue
		}
		for j := i * 2; j <= len(nums); j += i {
			primes[j] = true
		}
	}
	var a, b int
	for i := 0; i < len(nums); i++ {
		if primes[i] {
			a += nums[i]
		} else {
			b += nums[i]
		}
	}
	return int64(abs(a - b))
}
