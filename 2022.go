package main

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return nil
	}
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		ret[i] = original[start:end]
	}
	return ret
}
