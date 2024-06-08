package main

func superPow(a int, b []int) int {
	a %= 1337
	powHelper := func(m, k int) int {
		ret := 1
		for i := 0; i < k; i++ {
			ret *= m
			ret %= 1337
		}
		return ret
	}
	base := a
	ret := 1
	for i := len(b) - 1; i >= 0; i-- {
		ret *= powHelper(base, b[i])
		ret %= 1337
		base = powHelper(base, 10)
	}
	return ret
}
