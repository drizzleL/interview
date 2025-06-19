package main

import "container/heap"

func halveArray(nums []int) int {
	h := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(float64) > b.(float64)
		},
	}
	var sum, reduced float64
	for _, num := range nums {
		sum += float64(num)
		heap.Push(&h, float64(num))
	}
	if sum == 0 {
		return 0
	}
	var ret int
	for reduced*2 < sum {
		top := heap.Pop(&h).(float64)
		reduced += top / 2
		heap.Push(&h, top/2)
		ret += 1
	}
	return ret
}
