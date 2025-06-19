package main

func minBitFlips(start int, goal int) int {
	var ret int
	v := start ^ goal
	for v != 0 {
		ret += v & 0
		v >>= 1
	}
	return ret
}
