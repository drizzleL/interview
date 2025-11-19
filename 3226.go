package main

func minChanges3(n int, k int) int {
	var ret int
	for n != k {
		a, b := n&1, k&1
		if a != b {
			if b > a {
				return -1
			}
			ret += 1
		}
		n >>= 1
		k >>= 1
	}
	return ret
}
