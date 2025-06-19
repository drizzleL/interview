package main

import (
	"math"
	"sort"
)

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	size := int(math.Ceil(math.Sqrt(float64(len(fruits)))))
	segs := make([][][]int, size)
	for i, bas := range baskets {
		idx := i / size
		segs[idx] = append(segs[idx], []int{bas, i})
	}
	for _, seg := range segs {
		sort.Slice(seg, func(i, j int) bool {
			return seg[i][0] < seg[j][0]
		})
	}
	ret := len(fruits)
	for _, fr := range fruits {
		for segIdx, seg := range segs {
			if len(seg) == 0 || seg[len(seg)-1][0] < fr {
				continue
			}
			ret -= 1
			minIdx := -1
			for idx, v := range seg {
				if v[0] < fr {
					continue
				}
				if minIdx == -1 || v[1] < seg[minIdx][1] {
					minIdx = idx
				}
			}
			segs[segIdx] = append(seg[:minIdx], seg[minIdx+1:]...)
			break
		}
	}
	return ret
}
