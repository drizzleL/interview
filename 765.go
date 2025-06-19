package main

func minSwapsCouples(row []int) int {
	parent := make([]int, len(row))
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
	for i := 0; i < len(row)/2; i++ {
		union(row[i*2], row[i*2+1])
		union(i*2, i*2)
	}
	groups := map[int]int{}
	for i := range parent {
		groups[find(i)] += 1
	}
	var ret int
	for _, v := range groups {
		ret += v/2 - 1
	}
	return ret
}
