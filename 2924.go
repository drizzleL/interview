package main

func findChampion(n int, edges [][]int) int {
	in := make([]int, n)
	for _, e := range edges {
		in[e[1]] += 1
	}
	ret := -1
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			if ret != -1 {
				return -1
			}
			ret = i
		}
	}
	return ret
}
