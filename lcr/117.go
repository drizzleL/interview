package main

func numSimilarGroups(strs []string) int {
	dict := map[string]int{}
	var tt [][]int
	for m, str := range strs {
		if oldIdx, ok := dict[str]; ok {
			tt = append(tt, []int{oldIdx, m})
		} else {
			dict[str] = m
		}
		b := []byte(str)
		for i := 0; i < len(b); i++ {
			for j := i + 1; j < len(b); j++ {
				if b[i] == b[j] {
					continue
				}
				b[i], b[j] = b[j], b[i]
				if oldIdx, ok := dict[string(b)]; ok {
					tt = append(tt, []int{oldIdx, m})
				}
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	parent := make([]int, len(strs))
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(i, j int) {
		if i > j {
			i, j = j, i
		}
		parent[find(j)] = find(i)
	}
	for _, t := range tt {
		union(t[0], t[1])
	}
	var ret int
	for i, v := range parent {
		if i == v {
			ret++
		}
	}
	return ret
}
