package main

import (
	"math"
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func subarrayMajority(nums []int, queries [][]int) []int {
	blockSize := int(math.Sqrt(float64(len(nums))))
	for i, q := range queries {
		queries[i] = append(q, i)
	}
	sort.Slice(queries, func(i, j int) bool {
		q1, q2 := queries[i], queries[j]
		l1, l2 := q1[0], q2[0]
		b1, b2 := l1/blockSize, l2/blockSize
		if b1 == b2 {
			return q1[1] < q2[1]
		}
		return b1 < b2
	})
	rateCnt := map[int]int{}
	cnts := map[int]*redblacktree.Tree{}
	cntOrder := redblacktree.NewWithIntComparator()
	add := func(i int) {
		num := nums[i]
		rate := rateCnt[num]
		rateCnt[num] += 1
		if rate != 0 {
			oldTr := cnts[rate]
			oldTr.Remove(num)
			if oldTr.Empty() {
				cntOrder.Remove(rate)
			}
		}
		newTr := cnts[rate+1]
		if newTr == nil {
			newTr = redblacktree.NewWithIntComparator()
			cnts[rate+1] = newTr
		}
		newTr.Put(num, 0)
		cntOrder.Put(rate+1, 0)
	}
	del := func(i int) {
		num := nums[i]
		rate := rateCnt[num]
		rateCnt[num] -= 1
		oldTr := cnts[rate]
		oldTr.Remove(num)
		if oldTr.Empty() {
			cntOrder.Remove(rate)
		}
		if rate == 1 {
			return
		}
		newTr := cnts[rate-1]
		newTr.Put(num, 0)
		cntOrder.Put(rate-1, 0)
	}
	getMin := func(tr *redblacktree.Tree) int {
		val, _ := tr.Ceiling(math.MinInt32)
		return val.Key.(int)
	}
	getMax := func(tr *redblacktree.Tree) int {
		val, _ := tr.Floor(math.MaxInt32)
		return val.Key.(int)
	}
	getMaj := func() int {
		maxRate := getMax(cntOrder)
		tr := cnts[maxRate]
		return getMin(tr)
	}

	ret := make([]int, len(queries))
	currL, currR := 0, -1
	for _, q := range queries {
		l, r, t := q[0], q[1], q[2]
		for currR < r {
			currR += 1
			add(currR)
		}
		for currR > r {
			del(currR)
			currR -= 1
		}
		for currL > l {
			currL -= 1
			add(currL)
		}
		for currL < l {
			del(currL)
			currL += 1
		}
		maj := getMaj()
		if rateCnt[maj] < t {
			ret[q[3]] = -1
		} else {
			ret[q[3]] = maj
		}
	}
	return ret
}
