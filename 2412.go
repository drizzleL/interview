package main

import (
	"sort"
)

func minimumMoney(transactions [][]int) int64 {
	sort.Slice(transactions, func(i, j int) bool {
		t1, t2 := transactions[i], transactions[j]
		gap1, gap2 := t1[0]-t1[1], t2[0]-t2[1]
		if gap1 >= 0 && gap2 >= 0 {
			return t1[1] < t2[1]
		}
		if gap1 < 0 && gap2 < 0 {
			return t1[0] > t2[0]
		}
		return gap1 >= 0
	})
	var ret, now int
	for _, tr := range transactions {
		if tr[0] > now {
			ret += tr[0] - now
			now = tr[0]
		}
		now -= tr[0]
		now += tr[1]
	}
	return int64(ret)
}
