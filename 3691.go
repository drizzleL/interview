package main

import (
	"container/heap"
	"math"
)

func maxTotalValue2(nums []int, k int) int64 {
	size := len(nums)
	maxLog := int(math.Log2(float64(size)))
	st := &SparseTable{
		minData: make([][]int, size),
		maxData: make([][]int, size),
	}
	for i := 0; i < size; i++ {
		st.minData[i] = make([]int, maxLog+1)
		st.minData[i][0] = nums[i]
		st.maxData[i] = make([]int, maxLog+1)
		st.maxData[i][0] = nums[i]
	}
	for j := 1; j <= maxLog; j++ {
		for i := 0; i+(1<<j) <= size; i++ {
			st.minData[i][j] = min(st.minData[i][j-1], st.minData[i+(1<<(j-1))][j-1])
			st.maxData[i][j] = max(st.maxData[i][j-1], st.maxData[i+(1<<(j-1))][j-1])
		}
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([]int)[0] > b.([]int)[0]
		},
	}
	for i := 0; i < len(nums); i++ {
		heap.Push(h, []int{st.queryMax(0, i) - st.queryMin(0, i), 0, i})
	}
	var ret int
	for k > 0 {
		top := heap.Pop(h).([]int)
		ret += top[0]
		l, r := top[1], top[2]
		if l+1 <= r {
			heap.Push(h, []int{st.queryMax(l+1, r) - st.queryMin(l+1, r), l + 1, r})
		}
		k -= 1
	}
	return int64(ret)
}

type SparseTable struct {
	minData, maxData [][]int
}

func (st *SparseTable) queryLog(l, r int) int {
	size := r - l + 1
	return int(math.Log2(float64(size)))
}

func (st *SparseTable) queryMin(l, r int) int {
	j := st.queryLog(l, r)
	return min(st.minData[l][j], st.minData[r-1<<j+1][j])
}

func (st *SparseTable) queryMax(l, r int) int {
	j := st.queryLog(l, r)
	return max(st.maxData[l][j], st.maxData[r-1<<j+1][j])
}
