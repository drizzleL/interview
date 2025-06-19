package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func avoidFlood(rains []int) []int {
	rt := redblacktree.NewWithIntComparator()
	ret := make([]int, len(rains))
	dict := map[int]int{}
	for i, r := range rains {
		if r == 0 {
			ret[i] = 1
			rt.Put(i, i)
			continue
		}
		ret[i] = -1
		last, ok := dict[r]
		dict[r] = i
		if !ok {
			continue
		}
		node, ok := rt.Ceiling(last)
		if !ok {
			return nil
		}
		ret[node.Key.(int)] = r
		rt.Remove(node.Key)
	}
	return ret
}
