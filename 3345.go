package main

func smallestNumber(n int, t int) int {
	helper := func(x int) int {
		var ret int
		for x != 0 {
			ret *= x % 10
			x /= 10
		}
		return ret
	}
	for helper(n)%t != 0 {
		n += 1
	}
	return n
}
