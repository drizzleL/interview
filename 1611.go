package main

func minimumOneBitOperations(n int) int {
	ret := n
	for n >= 1 {
		n >>= 1
		ret ^= n
	}
	return ret
}
