package main

import (
	"math"
	"sort"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	ret := math.MinInt32
	for i1 := 0; i1 < m; i1++ {
		sums := make([]int, n)
		for i2 := i1; i2 < m; i2++ {
			for j := 0; j < n; j++ {
				sums[j] += matrix[i2][j]
			}
			ret = max(ret, subarrMaxSum(sums, k))
		}
	}
	return ret
}

func subarrMaxSum(nums []int, k int) int {
	ret := math.MinInt32
	ss := NewSortedSlice()
	ss.Insert(0)
	var presum int
	for _, num := range nums {
		presum += num
		val := ss.Ceiling(presum - k)
		if val != math.MinInt32 {
			ret = max(ret, presum-val)
		}
		ss.Insert(presum)
	}
	return ret
}

type SortedSlice struct {
	data []int
}

func NewSortedSlice() *SortedSlice {
	return &SortedSlice{}
}

func (s *SortedSlice) Insert(value int) {
	index := sort.SearchInts(s.data, value)
	s.data = append(s.data, 0) // 添加一个元素，以便在插入位置腾出空间
	copy(s.data[index+1:], s.data[index:])
	s.data[index] = value
}

func (s *SortedSlice) Ceiling(key int) int {
	index := sort.SearchInts(s.data, key)
	if index < len(s.data) {
		return s.data[index]
	}
	return math.MinInt32
}
