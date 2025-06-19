package main

import (
	"container/heap"
	"sort"
)

func findMaxSum(nums1 []int, nums2 []int, k int) []int64 {
	var eles [][3]int
	for i := range nums1 {
		eles = append(eles, [3]int{i, nums1[i], nums2[i]})
	}
	sort.Slice(eles, func(i, j int) bool {
		return eles[i][1] < eles[j][1]
	})
	var h IntHeap
	var sum int
	ret := make([]int64, len(nums1))
	for i := range eles {
		ele := eles[i]
		idx := ele[0]
		_, num2 := ele[1], ele[2]
		if i != 0 && eles[i][1] == eles[i-1][1] {
			ret[idx] = ret[eles[i-1][0]]
		} else {
			ret[idx] = int64(sum)
		}
		if h.Len() == k && h[0] < num2 {
			top := heap.Pop(&h).(int)
			sum -= top
		}
		if h.Len() < k {
			heap.Push(&h, num2)
			sum += num2
		}
	}
	return ret
}
