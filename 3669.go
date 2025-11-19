package main

import "math"

var primes []int

func init() {
	sieve := make([]bool, 100001)
	sieve[1] = true
	for i := 2; i < 100001; i++ {
		if sieve[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * 2; j < 100001; j += i {
			sieve[j] = true
		}
	}
}

func minDifference9(n int, k int) []int {
	ret := make([]int, k)
	minDiff := math.MaxInt32
	check := func(curr []int) {
		minVal, maxVal := math.MaxInt32, math.MinInt32
		for _, v := range curr {
			minVal = min(minVal, v)
			maxVal = max(maxVal, v)
		}
		if maxVal-minVal < minDiff {
			minDiff = maxVal - minVal
			copy(ret, curr)
		}
	}
	var dfs func(start int, n int, curr []int)
	dfs = func(start int, n int, curr []int) {
		if k-len(curr) == 1 {
			if n >= start {
				curr = append(curr, n)
				check(curr)
			}
			return
		}
		for i := start; i*i <= n; i++ {
			if n%i != 0 {
				continue
			}
			dfs(i, n/i, append(curr, i))
		}
	}
	dfs(1, n, nil)
	return ret
}
