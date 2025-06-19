package main

func countComponents(nums []int, threshold int) int {
	parent := make([]int, len(nums))
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
		if pa == pb {
			return
		}
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	dict := map[int]int{}
	for i, num := range nums {
		for j := 2; j*num <= threshold; j++ {
			k := num * j
			old, ok := dict[k]
			if ok {
				union(i, old)
			} else {
				dict[k] = i
			}
		}
	}
	group := map[int]bool{}
	for i := range parent {
		group[find(i)] = true
	}
	return len(group)
}
