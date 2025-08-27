package main

import "math/bits"

func minFlips9(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	ret := -1
	var num int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num |= mat[i][j] << (i*n + j)
		}
	}
	check := func(x int) bool {
		num2 := num
		var mask int
		for i := 0; i < m*n; i++ {
			if x&(1<<i) == 0 {
				continue
			}
			mask ^= 1 << i
			if i+1 < m*n {
				mask ^= 1 << (i + 1)
			}
			if i-1 >= 0 {
				mask ^= 1 << (i - 1)
			}
			if i+n < m*n {
				mask ^= 1 << (i + n)
			}
			if i-n >= 0 {
				mask ^= 1 << (i - n)
			}
		}
		return num2^mask == 0
	}
	for i := 0; i < 1<<(m*n); i++ {
		if !check(i) {
			continue
		}
		tmp := bits.OnesCount(uint(i))
		if ret == -1 || ret > tmp {
			ret = tmp
		}
	}
	return ret
}
