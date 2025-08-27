package main

func countHighestScoreNodes(parents []int) int {
	child := make([][]int, len(parents))
	for node, p := range parents {
		if p == -1 {
			continue
		}
		child[p] = append(child[p], node)
	}
	score := make([]int, len(parents))
	var dfs func(x int) int
	dfs = func(x int) int {
		ret := 1
		for _, c := range child[x] {
			ret += dfs(c)
		}
		score[x] = ret
		return ret
	}
	dfs(0)
	var ret, maxScore int
	for i := 0; i < len(parents); i++ {
		s := 1
		for _, c := range child[i] {
			s *= score[c]
		}
		if i != 0 {
			ret *= score[0] - score[i]
		}
		if s > maxScore {
			ret = 1
			maxScore = s
			continue
		}
		if s == maxScore {
			ret += 1
		}
	}
	return ret
}
