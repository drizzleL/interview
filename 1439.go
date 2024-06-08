package main

import (
	"container/heap"
	"strconv"
	"strings"
)

func kthSmallest(mat [][]int, k int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])
	visited := map[string]bool{}
	getKey := func(key []int) string {
		var ss []string
		for _, v := range key {
			ss = append(ss, strconv.Itoa(v))
		}
		return strings.Join(ss, "_")
	}
	var h KthItemPriorityQueue
	startItem := &KthItem{}
	for row := 0; row < m; row++ {
		startItem.key = append(startItem.key, 0)
		startItem.Sum += mat[row][0]
	}
	heap.Push(&h, startItem)
	visited[getKey(startItem.key)] = true
	for i := 1; i < k; i++ {
		top := heap.Pop(&h).(*KthItem)
		for row := 0; row < m; row++ {
			col := top.key[row]
			if col == n-1 {
				continue
			}
			top.key[row]++
			dictKey := getKey(top.key)
			if !visited[dictKey] {
				sum := top.Sum - mat[row][col] + mat[row][col+1]
				visited[dictKey] = true
				heap.Push(&h, &KthItem{
					key: append([]int{}, top.key...),
					Sum: sum,
				})
			}
			top.key[row]--
		}
	}

	return h[0].Sum
}

type KthItem struct {
	key []int
	Sum int
}

// A PriorityQueue implements heap.Interface and holds Items.
type KthItemPriorityQueue []*KthItem

func (pq KthItemPriorityQueue) Len() int { return len(pq) }

func (pq KthItemPriorityQueue) Less(i, j int) bool {
	return pq[i].Sum < pq[j].Sum
}

func (pq KthItemPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *KthItemPriorityQueue) Push(x any) {
	item := x.(*KthItem)
	*pq = append(*pq, item)
}

func (pq *KthItemPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
