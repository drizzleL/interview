package main

import "github.com/emirpasic/gods/trees/redblacktree"

func findXSum2(nums []int, k int, x int) []int64 {
	var ret []int64
	cmp := func(a, b interface{}) int {
		va, vb := a.(xsPair), b.(xsPair)
		if va.occ == vb.occ {
			return va.val - vb.val
		}
		return va.occ - vb.occ
	}
	xs := &xsum{
		dict:     map[int]int{},
		x:        x,
		largeSum: 0,
		large:    redblacktree.NewWith(cmp),
		small:    redblacktree.NewWith(cmp),
		cmp:      cmp,
	}
	for i := 0; i < len(nums); i++ {
		if i >= k {
			xs.Remove(nums[i-k])
		}
		xs.Add(nums[i])
		if i >= k-1 {
			ret = append(ret, int64(xs.largeSum))
		}
	}
	return ret
}

type xsum struct {
	dict         map[int]int
	x            int
	largeSum     int
	large, small *redblacktree.Tree
	cmp          func(a, b interface{}) int
}

type xsPair struct {
	val int
	occ int
}

func (xs *xsum) Remove(num int) {
	xs.remove(xsPair{num, xs.dict[num]})
	xs.dict[num] -= 1
	if xs.dict[num] > 0 {
		xs.add(xsPair{num, xs.dict[num]})
	}
}

func (xs *xsum) remove(p xsPair) {
	_, found := xs.large.Get(p)
	if found {
		xs.large.Remove(p)
		xs.largeSum -= p.val * p.occ
		if xs.small.Size() > 0 {
			topSmall := xs.small.Right().Key
			xs.small.Remove(topSmall)
			xs.large.Put(topSmall, nil)
			xs.largeSum += topSmall.(xsPair).val * topSmall.(xsPair).occ
		}
		return
	}
	xs.small.Remove(p)
}

func (xs *xsum) Add(num int) {
	if xs.dict[num] > 0 {
		xs.remove(xsPair{num, xs.dict[num]})
	}
	xs.dict[num] += 1
	xs.add(xsPair{num, xs.dict[num]})
}

func (xs *xsum) add(p xsPair) {
	if xs.large.Size() < xs.x {
		xs.largeSum += p.occ * p.val
		xs.large.Put(p, nil)
		return
	}
	minLarge := xs.large.Left().Key.(xsPair)
	if xs.cmp(p, minLarge) > 0 {
		xs.large.Remove(minLarge)
		xs.largeSum -= minLarge.val * minLarge.occ
		xs.large.Put(p, nil)
		xs.largeSum += p.val * p.occ
		xs.small.Put(minLarge, nil)
	} else {
		xs.small.Put(p, nil)
	}
}
