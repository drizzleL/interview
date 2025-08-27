package main

import "sort"

func processQueries2(c int, connections [][]int, queries [][]int) []int {
	parent := make([]int, c)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	for _, conn := range connections {
		union(conn[0]-1, conn[1]-1)
	}
	type net struct {
		nodes []int
		seen  map[int]bool
	}
	findNet := func(x *net, v int) int {
		if len(x.nodes) == 0 {
			return -1
		}
		if x.seen[v] { // online, return itselt
			return v + 1
		}
		return x.nodes[0] + 1
	}
	removeNet := func(x *net, v int) {
		x.seen[v] = false
		for len(x.nodes) > 0 && !x.seen[x.nodes[0]] {
			x.nodes = x.nodes[1:]
		}
	}
	dict := map[int]*net{}
	for i := range parent {
		p := find(i)
		x, ok := dict[p]
		if !ok {
			x = &net{
				seen: map[int]bool{},
			}
			dict[p] = x
		}
		x.nodes = append(x.nodes, i)
		x.seen[i] = true
	}
	for _, x := range dict {
		sort.Ints(x.nodes)
	}
	var ret []int
	for _, q := range queries {
		v := q[1] - 1
		x := dict[find(v)]
		if q[0] == 2 {
			removeNet(x, v)
			continue
		}
		ret = append(ret, findNet(x, v))
	}
	return ret
}
