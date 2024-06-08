package main

func multiply(A int, B int) int {
	var ret int
	if B > A {
		A, B = B, A
	}
	for B != 0 {
		if B&1 != 0 {
			ret += A
		}
		A += A
		B >>= 1
	}
	return ret
}
