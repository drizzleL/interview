package main

func networkDelayTime(times [][]int, n int, k int) int {
	sig := make([]int, n+1)
	for i := range sig {
		sig[i] = -1
	}

	dict := map[int][]int{}
	for i, item := range times {
		dict[item[0]] = append(dict[item[0]], i)
	}

	var q PriorityQueue
	q = append(q, &Item{
		cost: 0,
		idx:  k,
	})
	for len(q) != 0 {
		item := q.Pop().(*Item)
		sig[item.idx] = item.cost
		for _, nextIdx := range dict[item.idx] {
			next := times[nextIdx]
			newCost := item.cost + next[2]
			if sig[next[1]] != -1 {
				continue
			}
			q.Push(&Item{
				cost: newCost,
				idx:  next[1],
			})
		}
	}

	var ret int
	for i := 1; i < len(sig); i++ {
		v := sig[i]
		if v == -1 {
			return -1
		}
		ret = max(ret, v)
	}
	return ret
}

type Item struct {
	idx  int
	cost int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].idx < pq[j].idx
	}
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
