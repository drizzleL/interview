package main

func countPairs4(n int, edges [][]int) int64 {
	parent := make([]int, n)
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
	for _, ed := range edges {
		union(ed[0], ed[1])
	}
	dict := map[int]int{}
	for _, p := range parent {
		dict[find(p)] += 1
	}
	countHelper := func(x int) int {
		return (x * (x - 1)) / 2
	}
	ret := countHelper(n)
	for _, v := range dict {
		ret -= countHelper(v)
	}
	return int64(ret)
}
