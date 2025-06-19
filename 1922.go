package main

func countGoodNumbers(n int64) int {
	ret := 1
	if n%2 == 1 {
		ret *= 5
		n -= 1
	}
	pow := func(x, k int) int {
		ret := 1
		for k != 0 {
			if k&1 != 0 {
				ret *= x
				ret %= 1e9 + 7
			}
			k >>= 1
			x *= x
		}
		return ret
	}
	ret *= pow(20, int(n))
	ret %= 1e9 + 7
	return ret
}
