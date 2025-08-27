package main

import (
	"math/bits"
	"strconv"
)

var popCountComb [51][51]int
var popCountDict [51]int

func init() {
	for i := 2; i < len(popCountDict); i++ {
		popCountDict[i] = popCountDict[bits.OnesCount(uint(i))] + 1
	}
	popCountComb[0][0] = 1
	for i := 1; i < len(popCountComb); i++ {
		for j := 0; j <= i; j++ {
			popCountComb[i][j] = popCountComb[i-1][j]
			if j != 0 {
				popCountComb[i][j] += popCountComb[i-1][j-1]
			}
		}
	}
}

func popcountDepth(n int64, k int) int64 {
	if k == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	helper := func(bitSize int, preOneCnt int) int {
		var ret int
		for i := 0; i <= bitSize; i++ {
			if i+preOneCnt == 0 {
				continue
			}
			if popCountDict[i+preOneCnt]+1 != k {
				continue
			}
			ret += popCountComb[bitSize][i]
		}
		if preOneCnt == 0 && k == 1 {
			ret -= 1
		}
		return ret
	}
	oneCnt := bits.OnesCount(uint(n))
	var ret int
	if popCountDict[oneCnt]+1 == k {
		ret = 1
	}
	for i := 0; n != 0; i += 1 {
		if n&(1<<i) == 0 {
			continue
		}
		oneCnt -= 1
		ret += helper(i, oneCnt)
		n ^= 1 << i
	}
	return int64(ret)
}

func popcountDepth2(n int64, k int) int64 {
	if k == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	var oneCnt int
	str := strconv.FormatInt(n, 2)
	freq := make([]int, 51)
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			continue
		}
		leftSize := len(str) - i - 1
		for j := oneCnt; j < 51; j++ {
			freq[j] += popCountComb[leftSize][j-oneCnt]
		}
		oneCnt += 1
	}
	freq[oneCnt] += 1
	var ret int
	for i := 1; i < len(freq); i++ {
		if popCountDict[i]+1 != k {
			continue
		}
		ret += freq[i]
		if i == 1 && k == 1 {
			ret -= 1
		}
	}
	return int64(ret)
}
