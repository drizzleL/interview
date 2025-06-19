package main

func minEnd(n int, x int) int64 {
	ret := x
	n -= 1
	for base := 1; n != 0; base <<= 1 {
		if base&x != 0 {
			continue
		}
		if n&1 != 0 {
			ret |= base
		}
		n >>= 1
	}
	return int64(ret)
}
