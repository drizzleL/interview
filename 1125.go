package main

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	dict := map[string]int{}
	for _, ski := range req_skills {
		dict[ski] = len(dict)
	}
	dp := make([][]int, 1<<len(dict))
	dp[0] = []int{}
	for i, p := range people {
		var currSk int
		for _, sk := range p {
			idx := dict[sk]
			currSk |= 1 << idx
		}
		for prev := range dp {
			if dp[prev] == nil {
				continue
			}
			newSk := prev | currSk
			if newSk == prev {
				continue
			}
			if dp[newSk] == nil || len(dp[newSk]) > len(dp[prev])+1 {
				cp := make([]int, len(dp[prev]), len(dp[prev])+1)
				copy(cp, dp[prev])
				cp = append(cp, i)
				dp[newSk] = cp
			}
		}
	}
	return dp[(1<<len(dict))-1]
}
