package main

import (
	"container/heap"
	"math"
)

func minimumDifference2(nums []int) int64 {
	n := len(nums) / 3
	leftMin, rightMax := make([]int, len(nums)), make([]int, len(nums))
	h1 := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	var sum int
	for i := 0; i < n*2; i++ {
		sum += nums[i]
		heap.Push(h1, nums[i])
		if h1.Len() > n {
			sum -= heap.Pop(h1).(int)
		}
		leftMin[i] = sum
	}
	h2 := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) < b.(int)
		},
	}
	var sum2 int
	for i := len(nums) - 1; i >= n; i-- {
		sum2 += nums[i]
		heap.Push(h2, nums[i])
		if h2.Len() > n {
			sum2 -= heap.Pop(h2).(int)
		}
		rightMax[i] = sum2
	}
	ret := math.MaxInt64
	for i := n - 1; i < n*2; i++ {
		ret = min(ret, leftMin[i]-rightMax[i+1])
	}
	return int64(ret)
}
