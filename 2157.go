package main

func groupStrings(words []string) []int {
	parent := map[int]int{}
	for i := 0; i < len(words); i++ {
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
		if a == b {
			return
		}
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	wordToNum := func(x string) int {
		var ret int
		for _, c := range x {
			mask := 1 << int(c-'a')
			ret |= mask
		}
		return ret
	}
	dict := map[int]int{}
	check := func(i, num int) {
		if old, ok := dict[num]; ok {
			union(old, i)
		}
	}
	for m, w := range words {
		num := wordToNum(w)
		check(m, num)
		for i := 0; i < 26; i++ {
			if num&(1<<i) == 0 {
				check(m, num|(1<<i))
			} else {
				v := num & (^(1 << i))
				check(m, v)
				for j := 0; j < 26; j++ {
					if i == j {
						continue
					}
					if num&(1<<j) != 0 {
						continue
					}
					check(m, v|(1<<j))
				}
			}
		}
		dict[num] = m
	}
	groups := map[int][]int{}
	for i := range words {
		groups[find(i)] = append(groups[find(i)], i)
	}
	var maxSize int
	for _, g := range groups {
		maxSize = max(maxSize, len(g))
	}
	return []int{len(groups), maxSize}
}
