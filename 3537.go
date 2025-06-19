package main

import (
	"math"
)

func specialGrid(n int) [][]int {
	m := int(math.Pow(2, float64(n)))
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, m)
	}
	var helper func(up, bottom, left, right int, size int, l, r int)
	helper = func(up, bottom, left, right int, size int, l, r int) {
		if l == r {
			ret[up][left] = l
			return
		}
		valGap := size / 2 * size / 2
		helper(up, up+size/2-1, left+size/2, right, size/2, l, l+valGap-1)               // top-right
		helper(up+size/2, bottom, left+size/2, right, size/2, l+valGap, l+valGap*2-1)    // bot-right
		helper(up+size/2, bottom, left, left+size/2-1, size/2, l+valGap*2, l+valGap*3-1) // bot-left
		helper(up, up+size/2-1, left, left+size/2-1, size/2, l+valGap*3, r)              // top-left
	}
	helper(0, m-1, 0, m-1, m, 0, m*m-1)
	return ret
}
