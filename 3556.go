package main

import "sort"

func sumOfLargestPrimes(s string) int64 {
	isPrime := func(x int) bool {
		if x == 1 {
			return false
		}
		if x == 2 {
			return true
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	primes := map[int]bool{}
	for i := 0; i < len(s); i++ {
		var v int
		for j := i; j < len(s); j++ {
			d := int(s[j] - '0')
			v = v*10 + d
			if isPrime(v) {
				primes[v] = true
			}
		}
	}
	var pp []int
	for p := range primes {
		pp = append(pp, p)
	}
	sort.Ints(pp)
	var ret int
	for i, cnt := len(pp)-1, 0; i >= 0 && cnt < 3; i, cnt = i-1, cnt+1 {
		ret += pp[i]
	}
	return int64(ret)
}
