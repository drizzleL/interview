package main

import (
	"container/heap"
)

func getFinalState2(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	var h PriorityQueue
	var maxVal int
	for i, num := range nums {
		maxVal = max(maxVal, num)
		heap.Push(&h, &Item{
			cost: num,
			idx:  i,
		})
	}
	for ; k > 0; k-- {
		if h[0].cost*multiplier > maxVal {
			break
		}
		item := heap.Pop(&h).(*Item)
		item.cost *= multiplier
		item.cost %= 1e9 + 7
		heap.Push(&h, item)
		maxVal = max(maxVal, item.cost)
	}
	multi := func() int {
		n := k / len(nums)
		base := multiplier
		ret := 1
		for n != 0 {
			if n&1 != 0 {
				ret *= base
				ret %= 1e9 + 7
			}
			base *= base
			base %= 1e9 + 7
			n >>= 1
		}
		return ret
	}()
	for i := 0; i < k%len(nums); i++ {
		item := heap.Pop(&h).(*Item)
		item.cost *= multi
		item.cost %= 1e9 + 7
		item.cost *= multiplier
		item.cost %= 1e9 + 7
		nums[item.idx] = item.cost
	}
	for h.Len() != 0 {
		item := heap.Pop(&h).(*Item)
		item.cost *= multi
		item.cost %= 1e9 + 7
		nums[item.idx] = item.cost
	}
	return nums
}
