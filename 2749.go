package main

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	if num2 == 0 {
		return bits.OnesCount(uint(num1))
	}
	for k := 1; ; k++ {
		num1 -= num2
		if num1 < k {
			return -1
		}
		if num1 >= 0 && bits.OnesCount(uint(num1)) <= k {
			return k
		}
	}
	return -1
}
