package main

func makeIntegerBeautiful(n int64, target int) int64 {
	sum := func(x int64) int {
		var ret int
		for x > 0 {
			ret += int(x % 10)
			x /= 10
		}
		return ret
	}
	var ret int64
	for base := int64(10); sum(n) > target; base *= 10 {
		k := n % base
		ret += base - k
		n -= k
		n += base
	}
	return ret
}
