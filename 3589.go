package main

func init() {
	sieve := make([]bool, 1e5+1)
	sieve[0] = true
	sieve[1] = true
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		for j := i * 2; j < len(sieve); j += i {
			sieve[j] = true
		}
	}
}

var sieve []bool

func primeSubarray(nums []int, k int) int {
	var ret int
	var primes []int
	var minPrime, maxPrime []int
	canAdd := func(v int) bool {
		if sieve[v] { // not prime
			return true
		}
		if len(primes) == 0 {
			return true
		}
		maxVal := max(maxPrime[0], v)
		minVal := min(minPrime[0], v)
		return maxVal-minVal <= k
	}
	for i, j := 0, 0; i < len(nums); i++ {
		for ; j < len(nums) && canAdd(nums[j]); j++ {
			v := nums[j]
			if sieve[v] {
				continue
			}
			for len(minPrime) > 0 && minPrime[len(minPrime)-1] > v {
				minPrime = minPrime[:len(minPrime)-1]
			}
			minPrime = append(minPrime, v)
			for len(maxPrime) > 0 && maxPrime[len(maxPrime)-1] < v {
				maxPrime = maxPrime[:len(maxPrime)-1]
			}
			maxPrime = append(maxPrime, v)
			primes = append(primes, j)
		}
		if len(primes) >= 2 {
			ret += j - primes[1]
		}
		if sieve[nums[i]] {
			continue
		}
		primes = primes[1:]
		if len(minPrime) > 0 && minPrime[0] == nums[i] {
			minPrime = minPrime[1:]
		}
		if len(maxPrime) > 0 && maxPrime[0] == nums[i] {
			maxPrime = maxPrime[1:]
		}
	}
	return ret
}
