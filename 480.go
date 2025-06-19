package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	less, more := NewRbTree(), NewRbTree()
	balance := func() {
		for more.size > less.size {
			val := more.Min()
			cnt := more.RemoveAll(val)
			less.Add(val, cnt)
		}
		for less.size-more.size >= 2*less.MaxSize() {
			val := less.Max()
			cnt := less.RemoveAll(val)
			more.Add(val, cnt)
		}
	}
	getMedian := func() float64 {
		ret := float64(less.Max())
		if less.size == more.size {
			ret += float64(more.Min())
			ret /= 2
		}
		return ret
	}
	for i := 0; i < k; i++ {
		less.Add(nums[i], 1)
	}
	balance()
	ret := make([]float64, 0, len(nums)-k+1)
	ret = append(ret, getMedian())
	for i := k; i < len(nums); i++ {
		if nums[i] > less.Max() {
			more.Add(nums[i], 1)
		} else {
			less.Add(nums[i], 1)
		}
		prev := nums[i-k]
		if prev > less.Max() { // in more
			more.Remove(prev, 1)
		} else { // in less
			less.Remove(prev, 1)
		}
		balance()
		ret = append(ret, getMedian())
	}
	return ret
}

type RbTree struct {
	tr   *redblacktree.Tree
	size int
	sum  int
}

func NewRbTree() *RbTree {
	return &RbTree{
		tr: redblacktree.NewWithIntComparator(),
	}
}

func (m *RbTree) Max() int {
	node := m.tr.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Key.(int)
}

func (m *RbTree) MaxSize() int {
	node := m.tr.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Value.(int)
}

func (m *RbTree) Min() int {
	node := m.tr.Root
	for node.Left != nil {
		node = node.Left
	}
	return node.Key.(int)
}

func (m *RbTree) Add(val int, cnt int) {
	m.size += cnt
	m.sum += val * cnt
	node := m.tr.GetNode(val)
	if node != nil { // not found, just insert
		cnt += node.Value.(int)
	}
	m.tr.Put(val, cnt)
}

func (m *RbTree) RemoveAll(val int) int {
	node := m.tr.GetNode(val)
	cnt := node.Value.(int)
	m.Remove(val, cnt)
	return cnt
}

func (m *RbTree) Contains(val int) bool {
	return m.tr.GetNode(val) != nil
}

func (m *RbTree) Remove(val int, cnt int) {
	m.size -= cnt
	m.sum -= val * cnt
	nodeCnt := m.tr.GetNode(val).Value.(int)
	if nodeCnt == cnt {
		m.tr.Remove(val)
		return
	}
	m.tr.Put(val, nodeCnt-cnt)
}
