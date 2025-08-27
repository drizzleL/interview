package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minTime99(s string, order []int, k int) int {
	maxReach := (1 + len(s)) * len(s) / 2
	tr := redblacktree.NewWithIntComparator()
	if k > maxReach {
		return -1
	}
	tr.Put(-1, 0)
	tr.Put(len(s), 0)
	var valid int
	for i := 0; i < len(order); i++ {
		idx := order[i]
		before, _ := tr.Floor(idx)
		after, _ := tr.Ceiling(idx)
		fmt.Println(before.Key, after.Key)
		valid += (i - before.Key.(int)) * (after.Key.(int) - i)
		if valid >= k {
			return i
		}
		tr.Put(idx, 0)
	}
	return -1
}
