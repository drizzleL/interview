package main

func closestPrimes(left int, right int) []int {
	sieve := make([]bool, right+1)
	var primes []int
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		if i >= left && i <= right {
			primes = append(primes, i)
		}
		for j := 2; j*i < len(sieve); j++ {
			sieve[j*i] = true
		}
	}
	if len(primes) < 2 {
		return []int{-1, -1}
	}
	ret := []int{primes[0], primes[1]}
	for i := 2; i < len(primes); i++ {
		if primes[i]-primes[i-1] < ret[1]-ret[0] {
			ret[0] = primes[i-1]
			ret[1] = primes[i]
		}
	}
	return ret
}
