package main

import "sort"

func pondSizes(land [][]int) []int {
	if len(land) == 0 {
		return nil
	}
	m, n := len(land), len(land[0])
	vis := make([]bool, m*n)
	var traverse func(i, j int) int
	traverse = func(i, j int) int {
		if vis[i*n+j] {
			return 0
		}
		vis[i*n+j] = true
		if land[i][j] != 0 {
			return 0
		}
		ret := 1
		for _, d1 := range []int{0, -1, 1} {
			for _, d2 := range []int{0, -1, 1} {
				if d1 == 0 && d2 == 0 {
					continue
				}
				i2, j2 := d1+i, d2+j
				if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
					continue
				}
				ret += traverse(i2, j2)
			}
		}
		return ret
	}
	var ret []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] != 0 {
				continue
			}
			if vis[i*n+j] {
				continue
			}
			ret = append(ret, traverse(i, j))
		}
	}
	sort.Ints(ret)
	return ret
}
