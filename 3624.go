package main

import "math/bits"

func popcountDepth3(nums []int64, queries [][]int64) []int {
	pops := make([]int, len(nums))
	getVal := func(num int64) int {
		if num == 1 {
			return 0
		}
		return popCountDict[bits.OnesCount64(uint64(num))] + 1
	}
	for i, num := range nums {
		pops[i] = getVal(num)
	}
	bits := make([][]int, 6)
	for i := range bits {
		bits[i] = make([]int, len(nums)+1)
	}
	update := func(k int, i int64, diff int) {
		b := bits[k]
		i += 1
		for i < int64(len(b)) {
			b[i] += diff
			i += i & -i
		}
	}
	search := func(r, k int64) int {
		if r < 0 {
			return 0
		}
		r += 1
		b := bits[k]
		var ret int
		for r > 0 {
			ret += b[r]
			r -= r & -r
		}
		return ret
	}
	for i, pop := range pops {
		update(pop, int64(i), 1)
	}
	var ret []int
	for _, q := range queries {
		if q[0] == 2 { // for update
			oldVal, newVal := pops[q[1]], getVal(q[2])
			update(oldVal, q[1], -1)
			update(newVal, q[1], 1)
			pops[q[1]] = newVal
			continue
		}
		l, r, k := q[1], q[2], q[3]
		ret = append(ret, search(r, k)-search(l-1, k))
	}
	return ret
}
