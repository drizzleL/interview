package main

func minMaxWeight(n int, edges [][]int, threshold int) int {
	var l, r int
	for _, v := range edges {
		r = max(r, v[2])
	}
	dict := make([][]int, n)
	for i, ed := range edges {
		dict[ed[1]] = append(dict[ed[1]], i)
	}
	check := func(x int) bool {
		seen := make([]bool, n)
		seen[0] = true
		nodes := []int{0}
		for len(nodes) > 0 {
			var next []int
			for _, n := range nodes {
				for _, childIdx := range dict[n] {
					c := edges[childIdx]
					if c[2] > x {
						continue
					}
					if seen[c[0]] {
						continue
					}
					seen[c[0]] = true
					next = append(next, c[0])
				}
			}
			nodes = next
		}
		for _, v := range seen {
			if !v {
				return false
			}
		}
		return true
	}
	if !check(r) {
		return -1
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
