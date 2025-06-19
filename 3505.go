package main

import (
	"math"
)

type MedianRbTree struct {
	less, more *RbTree
}

func NewMedianRbTree() *MedianRbTree {
	return &MedianRbTree{
		less: NewRbTree(),
		more: NewRbTree(),
	}
}

func (mrb *MedianRbTree) Balance() {
	for mrb.more.size > mrb.less.size {
		val := mrb.more.Min()
		cnt := mrb.more.RemoveAll(val)
		mrb.less.Add(val, cnt)
	}
	for mrb.less.size-mrb.more.size >= 2*mrb.less.MaxSize() {
		val := mrb.less.Max()
		cnt := mrb.less.RemoveAll(val)
		mrb.more.Add(val, cnt)
	}
}

func (mrb *MedianRbTree) Median() float64 {
	ret := float64(mrb.less.Max())
	if mrb.less.size == mrb.more.size {
		ret += float64(mrb.more.Min())
		ret /= 2
	}
	return ret
}

func minOperations15(nums []int, x int, k int) int64 {
	mrb := NewMedianRbTree()
	getCost := func() int {
		median := mrb.less.Max()
		ret := median*mrb.less.size - mrb.less.sum
		ret += mrb.more.sum - median*mrb.more.size
		return ret
	}
	cost := make([]int, len(nums))
	for i := 0; i < x; i++ {
		mrb.less.Add(nums[i], 1)
	}
	mrb.Balance()
	cost[0] = getCost()
	for i := x; i < len(nums); i++ {
		if nums[i] > mrb.less.Max() {
			mrb.more.Add(nums[i], 1)
		} else {
			mrb.less.Add(nums[i], 1)
		}
		prev := nums[i-x]
		if prev > mrb.less.Max() { // in more
			mrb.more.Remove(prev, 1)
		} else { // in less
			mrb.less.Remove(prev, 1)
		}
		mrb.Balance()
		cost[i-x+1] = getCost()
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(nums))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
		for j := len(nums) - x; j >= 0; j-- {
			dp[0][j] = min(cost[j], dp[0][j+1])
		}
	}
	for i := 1; i < k; i++ {
		for j := len(nums) - (i+1)*x; j >= 0; j-- {
			dp[i][j] = dp[i][j+1]
			if j+x < len(nums) {
				dp[i][j] = min(dp[i][j], cost[j]+dp[i-1][j+x])
			}
		}
	}
	return int64(dp[k-1][0])
}
