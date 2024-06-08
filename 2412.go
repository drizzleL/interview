package main

import (
	"sort"
)

func minimumMoney(transactions [][]int) int64 {
	sort.Slice(transactions, func(i, j int) bool {
		t1, t2 := transactions[i], transactions[j]
		gap1, gap2 := t1[0]-t1[1], t2[0]-t2[1]
		if gap1 > 0 && gap2 > 0 {
			return gap1 > gap2
		}
		if gap1 > 0 {
			return true
		}
		if gap2 > 0 {
			return false
		}
		return t1[0] > t2[0]
	})
	last := []int{0, 0}
	for _, tran := range transactions {
		this := []int{0, 0}
		this[0] = max(last[0], (last[0]+tran[0])-min(last[1], tran[1]))
		this[1] = this[0] - last[0] - tran[0] + last[1] + tran[1]
		last = this
		if tran[0] <= tran[1] { // reach point
			break
		}
	}
	return int64(last[0])
}
