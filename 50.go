package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	ret := float64(1)
	for base := 1; base <= n; base *= 2 {
		if n&base != 0 {
			ret *= x
		}
		x *= x
	}
	return ret
}
