package main

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minAbsoluteDifference(nums []int, x int) int {
	rt := redblacktree.NewWithIntComparator()
	rt.Put(math.MaxInt32, nil)
	rt.Put(math.MinInt32, nil)
	ret := math.MaxInt32
	for i := x; i < len(nums); i++ {
		rt.Put(nums[i-x], nil)
		a, _ := rt.Ceiling(nums[i-x]) // find larger
		b, _ := rt.Floor(nums[i-x])   // find smaller
		ret = min(ret, a.Key.(int)-nums[i])
		ret = min(ret, nums[i]-b.Key.(int))
	}
	return ret
}
