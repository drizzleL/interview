package main

import "math"

func xorAfterQueries2(nums []int, queries [][]int) int {
	mod := int(1e9 + 7)
	b := int(math.Sqrt(float64(len(nums))))
	smallk := make([][][]int, b+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		k, v := q[2], q[3]
		if k > b {
			for i := l; i <= r; i += k {
				nums[i] *= v
				nums[i] %= mod
			}
			continue
		}
		smallk[k] = append(smallk[k], []int{l, r, v})
	}
	for k := 1; k <= b; k++ {
		if len(smallk[k]) == 0 {
			continue
		}
		diff := make([]int, len(nums))
		for i := range diff {
			diff[i] = 1
		}
		for _, v := range smallk[k] {
			l, r, v := v[0], v[1], v[2]
			diff[l] *= v
			diff[l] %= mod
			r2 := ((r-l)/k+1)*k + l
			if r2 < len(nums) {
				diff[r2] *= pow(v, mod-2, mod)
				diff[r2] %= mod
			}
		}
		for start := 0; start < k; start++ {
			curr := 1
			for i := start; i < len(nums); i += k {
				curr *= diff[i]
				curr %= mod
				nums[i] *= curr
				nums[i] %= mod
			}
		}
	}
	var ret int
	for _, num := range nums {
		ret ^= num
	}
	return ret
}
