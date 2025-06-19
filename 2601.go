package main

import "sort"

func primeSubOperation(nums []int) bool {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]bool, maxVal+1)
	var primes []int
	for i := 2; i < len(sieve); i++ {
		if sieve[i] { // even pass
			continue
		}
		primes = append(primes, i)
		for j := 2; j*i < len(sieve); j++ {
			sieve[i*j] = true
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			continue
		}
		j := sort.SearchInts(primes, nums[i]-nums[i+1]+1)
		if j == len(primes) || primes[j] >= nums[i] {
			return false
		}
		nums[i] -= primes[j]
	}
	return true
}
