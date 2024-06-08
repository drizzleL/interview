package main

func findJudge(n int, trust [][]int) int {
	dict := make([]int, n)
	ever := make([]bool, n)
	for _, t := range trust {
		dict[t[1]-1]++
		ever[t[0]-1] = true
	}
	for k, v := range dict {
		if v == n-1 && !ever[k] {
			return k + 1
		}
	}
	return -1
}
