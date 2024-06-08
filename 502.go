package main

import (
	"container/heap"
	"sort"
)

type IpoItem struct {
	capital int
	profit  int
}

type IpoItems []*IpoItem

func (pq IpoItems) Len() int { return len(pq) }

func (pq IpoItems) Less(i, j int) bool {
	return pq[i].profit > pq[j].profit
}

func (pq IpoItems) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IpoItems) Push(x any) {
	item := x.(*IpoItem)
	*pq = append(*pq, item)
}

func (pq *IpoItems) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var items IpoItems
	for i := 0; i < len(profits); i++ {
		items = append(items, &IpoItem{profit: profits[i], capital: capital[i]})
	}
	// first sort by capital
	sort.Slice(items, func(i, j int) bool {
		return items[i].capital < items[j].capital
	})
	var j int

	var profitItems IpoItems
	for i := 0; i < k; i++ {
		for j < len(items) && items[j].capital <= w {
			heap.Push(&profitItems, items[j])
			j++
		}
		if len(profitItems) == 0 {
			break
		}
		item := heap.Pop(&profitItems).(*IpoItem)
		w += item.profit
	}
	return w
}
