package main

import (
	"math"
	"sort"
)

func numOfUnplacedFruits3(fruits []int, baskets []int) int {
	segTree := make([]int, len(baskets)*4)
	var build func(i int, l, r int)
	build = func(i int, l, r int) {
		if l == r {
			segTree[i] = baskets[l]
			return
		}
		mid := l + (r-l)/2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		segTree[i] = max(segTree[i*2], segTree[i*2+1])
	}
	build(1, 0, len(baskets)-1)
	var query func(i int, l, r int, val int) (idx int)
	query = func(i int, l, r int, val int) (idx int) {
		if segTree[i] < val {
			return -1
		}
		if l == r {
			return l
		}
		mid := l + (r-l)/2
		idx = query(i*2, l, mid, val)
		if idx != -1 {
			return
		}
		return query(i*2+1, mid+1, r, val)
	}
	var update func(i int, l, r, idx, val int)
	update = func(i int, l, r, idx, val int) {
		if l == r {
			segTree[i] = val
			return
		}
		mid := l + (r-l)/2
		if idx <= mid {
			update(i*2, l, mid, idx, val)
		} else {
			update(i*2+1, mid+1, r, idx, val)
		}
		segTree[i] = max(segTree[i*2], segTree[i*2+1])
	}
	var ret int
	for _, f := range fruits {
		idx := query(1, 0, len(baskets)-1, f)
		if idx == -1 {
			ret += 1
			continue
		}
		update(1, 0, len(baskets)-1, idx, 0)
	}
	return ret
}
func numOfUnplacedFruits(fruits []int, baskets []int) int {
	size := int(math.Ceil(math.Sqrt(float64(len(baskets)))))
	var segs [][][]int
	for i, bas := range baskets {
		idx := i / size
		if idx >= len(segs) {
			segs = append(segs, make([][]int, 0))
		}
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
