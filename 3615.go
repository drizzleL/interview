package main

func maxLen(n int, edges [][]int, label string) int {
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, 1<<n)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	var helper func(a, b int, mask int) int
	helper = func(a, b int, mask int) (ret int) {
		if cache[a][b][mask] != -1 {
			return cache[a][b][mask]
		}
		defer func() {
			cache[a][b][mask] = ret
		}()
		if label[a] != label[b] {
			return 0
		}
		for _, next := range dict[b] {
			if (1<<next)&mask != 0 {
				continue
			}
			mask2 := mask | 1<<next
			for _, next2 := range dict[a] {
				if (1<<next2)&mask2 != 0 {
					continue
				}
				if label[next] != label[next2] {
					continue
				}
				ret = max(ret, helper(next, next2, mask2|(1<<next2))+2)
			}
		}
		return ret
	}
	var ret int
	for i := 0; i < n; i++ {
		ret = max(ret, helper(i, i, 1<<i)+1)
	}
	for _, ed := range edges {
		if label[ed[0]] != label[ed[1]] {
			continue
		}
		ret = max(ret, helper(ed[0], ed[1], (1<<ed[0])|(1<<ed[1]))+2)
	}
	return ret
}
