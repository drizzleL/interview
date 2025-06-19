package main

import (
	"math"
)

func crackSafe(n int, k int) string {
	size := int(math.Pow(float64(k), float64(n)))
	seen := make([]bool, size)
	var ret string
	var helper func(b []byte, num, cnt int) bool
	helper = func(b []byte, num, cnt int) bool {
		if cnt == len(seen) {
			ret = string(b)
			return true
		}
		num = (num * k) % size
		for i := 0; i < k; i++ {
			if seen[num+i] {
				continue
			}
			seen[num+i] = true
			if helper(append(b, byte('0'+i)), num+i, cnt+1) {
				return true
			}
			seen[num+i] = false
		}
		return false
	}
	b := make([]byte, n, len(seen)+n-1)
	for i := range b {
		b[i] = '0'
	}
	seen[0] = true
	helper(b, 0, 1)
	return ret
}
