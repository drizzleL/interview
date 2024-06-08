package main

import (
	"container/heap"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	nonNegDict := map[int][][2]int{}
	for i, ed := range edges {
		if ed[2] == -1 {
			continue
		}
		nonNegDict[ed[0]] = append(nonNegDict[ed[0]], [2]int{ed[1], i})
		nonNegDict[ed[1]] = append(nonNegDict[ed[1]], [2]int{ed[0], i})
	}
	dict := map[int][][2]int{}
	for i, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], i})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], i})
	}
	check := func(dict map[int][][2]int) int {
		var retNode *TopoItem
		var items TopoItems
		heap.Push(&items, &TopoItem{
			Node: source,
			Cost: 0,
		})
		var flag bool
		visited := make([]bool, n)
		visited2 := make([]bool, n)
		for items.Len() != 0 {
			item := heap.Pop(&items).(*TopoItem)
			if len(item.Idxs) == 0 {
				if visited[item.Node] {
					continue
				}
				visited[item.Node] = true
			} else {
				if visited2[item.Node] {
					continue
				}
				visited2[item.Node] = true
			}
			if item.Node == destination {
				if len(item.Idxs) == 0 && item.Cost != target {
					return 1
				}
				flag = true
				if len(item.Idxs) != 0 {
					retNode = item
				}
				continue
			}
			for _, next := range dict[item.Node] {
				cost := edges[next[1]][2]
				if cost > 0 {
					if item.Cost+cost > target {
						continue
					} else {
						heap.Push(&items, &TopoItem{
							Node: next[0],
							Cost: item.Cost + cost,
							Idxs: item.Idxs,
						})
					}
				} else {
					topoItem := &TopoItem{
						Node: next[0],
						Cost: item.Cost + 1,
						Idxs: append(item.Idxs, next[1]),
					}
					heap.Push(&items, topoItem)
				}
			}
		}
		if !flag {
			return 0
		}
		if retNode != nil {
			edges[retNode.Idxs[0]][2] = target - retNode.Cost + 1
			for i := 1; i < len(retNode.Idxs); i++ {
				edges[retNode.Idxs[i]][2] = 1
			}
		}
		return 2
	}
	if check(nonNegDict) == 1 {
		return nil
	}
	if check(dict) != 2 {
		return nil
	}
	for i, ed := range edges {
		if ed[2] == -1 {
			edges[i][2] = target
		}
	}
	return edges
}

type TopoItem struct {
	Node int
	Cost int
	Idxs []int
}

type TopoItems []*TopoItem

func (h TopoItems) Len() int           { return len(h) }
func (h TopoItems) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h TopoItems) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TopoItems) Push(x any) {
	*h = append(*h, x.(*TopoItem))
}

func (h *TopoItems) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
