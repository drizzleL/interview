package main

func findPrimePairs(n int) [][]int {
	primes := make([]bool, n/2+1)
	for i := 2; i < n; i++ {
		if primes[i] == true {
			continue
		}
		for j := 2; i*j < len(primes); j++ {
			primes[i*j] = true
		}
	}
	var ret [][]int
	for i := 2; i <= n/2; i++ {
		if primes[i] || primes[n-i] {
			continue
		}
		ret = append(ret, []int{i, n - i})
	}
	return ret
}
