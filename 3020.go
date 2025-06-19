package main

import (
	"math"
	"sort"
)

func maximumLength2(nums []int) int {
	var oneCnt int
	sort.Ints(nums)
	dict := map[int]int{}
	streak := map[int]int{}
	ret := 1
	for _, num := range nums {
		if num == 1 {
			oneCnt += 1
			continue
		}
		dict[num] += 1
		streak[num] = 1
		prev := int(math.Sqrt(float64(num)))
		if prev*prev != num { // no right sqrt
			continue
		}
		if dict[prev] < 2 { // no enough prev
			continue
		}
		streak[num] = streak[prev] + 1
		ret = max(ret, streak[num]*2-1)
	}
	ret = max(ret, (oneCnt-1)/2*2+1)
	return ret
}
