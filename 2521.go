package main

func distinctPrimeFactors(nums []int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]bool, maxVal+1)
	primes := map[int]bool{}
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		primes[i] = true
		for j := 2; j*i < len(sieve); j++ {
			sieve[j*i] = true
		}
	}
	var ret int
	for _, num := range nums {
		var picked []int
		for p := range primes {
			if p > num {
				continue
			}
			if num%p == 0 {
				picked = append(picked, p)
			}
		}
		ret += len(picked)
		for _, p := range picked {
			delete(primes, p)
		}
	}
	return ret
}
