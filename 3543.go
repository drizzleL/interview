package main

func maxWeight2(n int, edges [][]int, k int, t int) int {
	dict := make([][][2]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
	}
	cache := map[[3]int][]int{}
	var helper func(i int, k int) []int
	helper = func(i int, k int) (ret []int) {
		if k == 0 {
			return []int{0}
		}
		if c, ok := cache[[3]int{i, k}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{i, k}] = ret
		}()
		seen := map[int]bool{}
		for _, next := range dict[i] {
			j, w := next[0], next[1]
			for _, next := range helper(j, k-1) {
				w2 := w + next
				if w2 > t || seen[w2] {
					continue
				}
				seen[w2] = true
				ret = append(ret, w2)
			}
		}
		return
	}
	ret := -1
	for i := 0; i < n; i++ {
		for _, sum := range helper(i, k) {
			ret = max(ret, sum)
		}
	}
	return ret
}
