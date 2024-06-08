package main

func edgeScore(edges []int) int {
	var ret int
	scores := make([]int, len(edges))
	for i, ed := range edges {
		scores[ed] += i
		if (scores[ed] > scores[ret]) || (scores[ed] == scores[ret] && ed < ret) {
			ret = ed
		}
	}
	return ret
}
