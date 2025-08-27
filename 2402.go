package main

import (
	"container/heap"
	"sort"
)

func mostBooked2(n int, meetings [][]int) int {
	f := func(a, b interface{}) bool {
		av, bv := a.([]int), b.([]int)
		if av[0] == bv[0] {
			return av[1] < bv[1]
		}
		return av[0] < bv[0]
	}
	used := &HeapArr{
		LessHelper: f,
	}
	free := &HeapArr{
		LessHelper: f,
	}
	for i := 0; i < n; i++ {
		heap.Push(free, []int{0, i})
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	dict := make([]int, n)
	for _, met := range meetings {
		start, end := met[0], met[1]
		for used.Len() != 0 && used.Top().([]int)[0] <= start {
			top := heap.Pop(used).([]int)
			heap.Push(free, []int{0, top[1]})
		}
		if free.Len() == 0 { // pop first used
			top := heap.Pop(used).([]int)
			heap.Push(free, []int{0, top[1]})
			end = top[0] + end - start
		}
		top := heap.Pop(free).([]int)
		heap.Push(used, []int{end, top[1]})
		dict[top[1]]++
	}
	var ret int
	for k, v := range dict {
		if v > dict[ret] {
			ret = k
		}
	}
	return ret + 1
}
func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var free MeetingItems
	var used MeetingItems
	dict := make([]int, n)
	for i := 0; i < n; i++ {
		heap.Push(&free, &MeetingItem{
			start:  0,
			number: i,
		})
	}
	for _, m := range meetings {
		for used.Len() != 0 && used[0].start <= m[0] {
			item := heap.Pop(&used).(*MeetingItem)
			item.start = 0 // init as 0
			heap.Push(&free, item)
		}
		if free.Len() != 0 {
			item := heap.Pop(&free).(*MeetingItem)
			item.start = m[1]
			heap.Push(&used, item)
			dict[item.number]++
			continue
		}
		item := heap.Pop(&used).(*MeetingItem)
		item.start = max(item.start, m[0]) + m[1] - m[0]
		heap.Push(&used, item)
		dict[item.number]++
	}
	var retIdx int
	for i, v := range dict {
		if v > dict[retIdx] {
			retIdx = i
		}
	}
	return retIdx
}

type MeetingItem struct {
	start  int
	number int
}

type MeetingItems []*MeetingItem

func (pq MeetingItems) Len() int { return len(pq) }

func (pq MeetingItems) Less(i, j int) bool {
	if pq[i].start == pq[j].start {
		return pq[i].number < pq[j].number
	}
	return pq[i].start < pq[j].start
}

func (pq MeetingItems) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MeetingItems) Push(x any) {
	item := x.(*MeetingItem)
	*pq = append(*pq, item)
}

func (pq *MeetingItems) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
