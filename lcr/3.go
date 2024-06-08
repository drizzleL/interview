package main

func countBits(n int) []int {
	contain := func(v int) int {
		var ret int
		for v != 0 {
			ret += v & 1
			v >>= 1
		}
		return ret
	}
	ret := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ret[i] = contain(i)
	}
	return ret
}
