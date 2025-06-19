package main

func smallestNumber3(n int) int {
	base := 1
	for n >= base {
		n |= base
		base <<= 1
	}
	return n
}
