package main

import (
	"math/bits"
	"sort"
)

func subarrayBitwiseORs(arr []int) int {
	var bitSize int
	for _, v := range arr {
		bitSize = max(bitSize, bits.Len(uint(v)))
	}
	dict := make([][]int, bitSize)
	for i, num := range arr {
		for j := 0; j < bitSize; j++ {
			if (num>>j)&1 == 0 {
				continue
			}
			dict[j] = append(dict[j], i)
		}
	}
	seen := map[int]bool{}
	cnts := make([]int, bitSize)
	for _, num := range arr {
		seen[num] = true
		var tmp [][2]int
		for j := 0; j < bitSize; j++ {
			if (num>>j)&1 != 0 { // already seen
				cnts[j] += 1
				continue
			}
			if len(dict[j]) == cnts[j] { // after no exists
				continue
			}
			tmp = append(tmp, [2]int{j, dict[j][cnts[j]]})
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i][1] < tmp[j][1]
		})
		for i, v := range tmp {
			j := v[0]
			num |= 1 << j
			if i == len(tmp)-1 || v[1] != tmp[i+1][1] {
				seen[num] = true
			}
		}
	}
	return len(seen)
}
