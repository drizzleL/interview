package main

import (
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	last := 0
	var i int
	var v int
	for ; i < len(s); i++ {
		curr := int(s[i] - '0')
		if curr < last {
			break
		}
		v = v*10 + curr
		last = curr
	}
	if i == len(s) {
		return n
	}
	for {
		i--
		if i == 0 {
			break
		}
		if s[i-1] != s[i] {
			break
		}
		v /= 10
	}
	base := 1
	for ; i < len(s)-1; i++ {
		base *= 10
	}
	return v*base - 1
}
