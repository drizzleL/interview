package main

func countGoodArrays(n int, m int, k int) int {
	ret := fastPow(m-1, n-1-k, m)
	ret *= combination(n-1, k)
	ret %= 1e9 + 7
	return ret
}
