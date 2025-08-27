package main

import (
	"math"
	"sort"
)

func maxJumps(arr []int, d int) int {
	left, right := make([][]int, len(arr)), make([][]int, len(arr))
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && i-j <= d; j-- {
			if arr[j] >= arr[i] {
				break
			}
			if len(left[i]) == 0 {
				left[i] = append(left[i], j)
				continue
			}
			if arr[j] < arr[left[i][0]] {
				continue
			}
			if arr[j] > arr[left[i][0]] {
				left[i] = left[i][:0]
			}
			left[i] = append(left[i], j)
		}
	}
	for i := len(arr) - 2; i >= 0; i-- {
		for j := i + 1; j < len(arr) && j-i <= d; j++ {
			if arr[j] >= arr[i] {
				break
			}
			if len(right[i]) == 0 {
				right[i] = append(right[i], j)
				continue
			}
			if arr[j] < arr[right[i][0]] {
				continue
			}
			if arr[j] > arr[right[i][0]] {
				right[i] = right[i][:0]
			}
			right[i] = append(right[i], j)
		}
	}
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = -1
	}
	var helper func(i int) int
	helper = func(i int) int {
		if dp[i] != -1 {
			return dp[i]
		}
		dp[i] = 1
		for _, v := range left[i] {
			dp[i] = max(dp[i], helper(v)+1)
		}
		for _, v := range right[i] {
			dp[i] = max(dp[i], helper(v)+1)
		}
		return dp[i]
	}
	var ret int
	for i := 0; i < len(arr); i++ {
		ret = max(ret, helper(i))
	}
	return ret
}

func maxJumps2(arr []int, d int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}
	type node struct {
		prev, next *node
		val        int
		idx        int
	}
	eles := make([]*node, 0, len(arr))
	var prev *node
	for i, v := range arr {
		ele := &node{
			val: v,
			idx: i,
		}
		if prev != nil {
			prev.next = ele
			ele.prev = prev
		}
		eles = append(eles, ele)
		prev = ele
	}
	sort.Slice(eles, func(i, j int) bool {
		if eles[i].val == eles[j].val {
			return eles[i].idx < eles[j].idx
		}
		return eles[i].val < eles[j].val
	})
	helper := func(eles []*node) {
		var validPrev, validNext *node
		for i := 0; i < len(eles); i++ {
			prev := eles[i].prev
			if prev == nil || prev.val == eles[i].val {
				prev = validPrev
			}
			if prev != nil {
				if eles[i].idx-prev.idx <= d {
					dp[prev.idx] = max(dp[prev.idx], dp[eles[i].idx]+1)
				}
			}
			validPrev = prev
		}
		for i := len(eles) - 1; i >= 0; i-- {
			next := eles[i].next
			if next == nil || next.val == eles[i].val {
				next = validNext
			}
			if next != nil {
				if next.idx-eles[i].idx <= d {
					dp[next.idx] = max(dp[next.idx], dp[eles[i].idx]+1)
				}
			}
			validNext = next
		}
		for _, ele := range eles {
			prev, next := ele.prev, ele.next
			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.prev = prev
			}
		}
	}
	for i := 0; i < len(eles); {
		var s []*node
		v := eles[i].val
		for ; i < len(eles) && eles[i].val == v; i++ {
			s = append(s, eles[i])
		}
		helper(s)
	}
	var ret int
	for _, v := range dp {
		ret = max(ret, v)
	}
	return ret
}

func maxJumps3(arr []int, d int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}
	var stack []int
	arr = append(arr, math.MaxInt32)
	for i := 0; i < len(arr); i++ {
		for len(stack) != 0 && arr[stack[len(stack)-1]] < arr[i] {
			var l []int
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l = append(l, j)
			for len(stack) != 0 && arr[stack[len(stack)-1]] == arr[j] {
				j2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				l = append(l, j2)
			}
			for _, j := range l {
				if i-j <= d {
					dp[i] = max(dp[i], dp[j]+1)
				}
				if len(stack) != 0 && j-stack[len(stack)-1] <= d {
					k := stack[len(stack)-1]
					dp[k] = max(dp[k], dp[j]+1)
				}
			}
		}
		stack = append(stack, i)
	}
	var ret int
	for _, v := range dp {
		ret = max(ret, v)
	}
	return ret
}
