package main

import (
	"container/heap"
	"sort"
)

func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([]int)[1] > b.([]int)[1]
		},
	}
	diff := make([]int, len(nums)+1)
	var cnts int
	ret := len(queries)
	for i, j := 0, 0; i < len(nums); i++ {
		cnts += diff[i]
		if cnts >= nums[i] {
			continue
		}
		for ; j < len(queries) && queries[j][0] <= i; j++ {
			if queries[j][1] < i { // useless, just remove
				continue
			}
			heap.Push(h, queries[j])
		}
		for cnts < nums[i] && h.Len() != 0 {
			top := heap.Pop(h).([]int)
			if top[1] < i { // useless
				continue
			}
			diff[top[1]+1] -= 1
			cnts += 1
			ret -= 1
		}
		if cnts < nums[i] { // still not match, just return
			return -1
		}
	}
	return ret
}
