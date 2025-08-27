package main

import (
	"container/heap"
)

func minJumps9(nums []int) int {
	dict := map[int][]int{}
	var maxVal int
	for i, num := range nums {
		dict[num] = append(dict[num], i)
		maxVal = max(maxVal, num)
	}
	jumpDict := make([][]int, len(nums))
	sieve := make([]bool, maxVal+1)
	sieve[1] = true
	for i := 2; i <= maxVal; i++ {
		if sieve[i] {
			continue
		}
		for j := i * 2; j <= maxVal; j += i {
			sieve[j] = true
			for _, start := range dict[i] {
				for _, end := range dict[j] {
					jumpDict[start] = append(jumpDict[start], end)
				}
			}
		}
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[0] < b.([2]int)[0]
		},
	}
	heap.Push(h, [2]int{0, 0})
	seen := make([]int, len(nums))
	for i := 1; i < len(seen); i++ {
		seen[i] = -1
	}
	for h.Len() > 0 {
		top := heap.Pop(h).([2]int)
		cnt, i := top[0], top[1]
		if i == len(nums)-1 {
			return cnt
		}
		if i-1 >= 0 && seen[i-1] == -1 {
			seen[i-1] = cnt + 1
			heap.Push(h, [2]int{cnt + 1, i - 1})
		}
		if i+1 < len(nums) && seen[i+1] == -1 {
			seen[i+1] = cnt + 1
			heap.Push(h, [2]int{cnt + 1, i + 1})
		}
		for _, next := range jumpDict[i] {
			if seen[next] != -1 {
				continue
			}
			seen[next] = cnt + 1
			heap.Push(h, [2]int{cnt + 1, next})
		}
		if sieve[nums[i]] {
			continue
		}
		for _, next := range dict[nums[i]] {
			if seen[next] != -1 {
				continue
			}
			seen[next] = cnt + 1
			heap.Push(h, [2]int{cnt + 1, next})
		}
	}
	return 0
}
