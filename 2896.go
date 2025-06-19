package main

import "math"

func minOperations6(s1 string, s2 string, x int) int {
	var ret, odd int
	pre := math.MaxInt32
	for i := range s1 {
		pre += 2
		if s1[i] == s2[i] {
			continue
		}
		odd ^= 1
		last := ret
		ret = min(ret+x*2, pre)
		pre = last
	}
	if odd != 0 {
		return -1
	}
	return ret / 2
}
