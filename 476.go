package main

func findComplement(num int) int {
	base := 1
	var ret int
	for num > 0 {
		if num&1 == 0 {
			ret |= base
		}
		base <<= 1
		num >>= 1
	}
	return ret
}
