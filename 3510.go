package main

import "container/heap"

func minimumPairRemoval(nums []int) int {
	type item struct {
		sum     int
		leftIdx int
		closed  bool
	}
	dict := make([]*item, len(nums))
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			va, vb := a.(*item), b.(*item)
			if va.sum == vb.sum {
				return va.leftIdx < vb.leftIdx
			}
			return va.sum < vb.sum
		},
	}
	prev, next := make([]int, len(nums)), make([]int, len(nums))
	var cnt int
	for i := 0; i < len(nums); i++ {
		prev[i] = i - 1
		next[i] = i + 1
		if i == len(nums)-1 {
			continue
		}
		if nums[i+1] < nums[i] {
			cnt += 1
		}
		v := &item{
			sum:     nums[i+1] + nums[i],
			leftIdx: i,
		}
		dict[i] = v
		heap.Push(h, v)
	}
	var ret int
	for cnt != 0 {
		top := heap.Pop(h).(*item)
		if top.closed {
			continue
		}
		i, j := top.leftIdx, next[top.leftIdx]
		if nums[i] > nums[j] {
			cnt -= 1
		}
		p, q := prev[i], next[j]
		if p >= 0 {
			lItem := dict[p]
			lItem.closed = true
			v := &item{
				sum:     nums[p] + top.sum,
				leftIdx: p,
			}
			dict[p] = v
			heap.Push(h, v)
			if nums[p] > nums[i] && nums[p] <= top.sum {
				cnt -= 1
			}
			if nums[p] <= nums[i] && nums[p] > top.sum {
				cnt += 1
			}
			next[p] = i
		}
		prev[i] = p
		if q < len(nums) {
			rItem := dict[j]
			rItem.closed = true
			v := &item{
				sum:     nums[q] + top.sum,
				leftIdx: i,
			}
			dict[i] = v
			heap.Push(h, v)
			if nums[j] <= nums[q] && top.sum > nums[q] {
				cnt += 1
			}
			if nums[j] > nums[q] && top.sum <= nums[q] {
				cnt -= 1
			}
			prev[q] = i
		}
		next[i] = q
		nums[i] = top.sum
		ret += 1
	}
	return ret
}
