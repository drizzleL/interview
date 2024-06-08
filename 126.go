package main

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	endIdx := -1
	group := map[string][]int{}
	for i, w := range wordList {
		if w == endWord {
			endIdx = i
		}
		for j := 0; j < len(w); j++ {
			key := w[:j] + "#" + w[j+1:]
			group[key] = append(group[key], i)
		}
	}
	if endIdx == -1 { // nof found end word
		return nil
	}
	nextGroup := map[int][]int{}
	for _, g := range group {
		for i := 0; i < len(g); i++ {
			for j := 0; j < len(g); j++ {
				if i == j {
					continue
				}
				nextGroup[g[i]] = append(nextGroup[g[i]], g[j])
			}
		}
	}

	similar := func(a, b string) bool {
		var diff int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff += 1
			}
		}
		return diff == 1
	}

	seen := make([]bool, len(wordList))
	var helper func(start int, curr []int)
	var ret [][]string
	maxSize := math.MaxInt32
	helper = func(start int, curr []int) {
		if start == endIdx {
			if len(curr) < maxSize {
				maxSize = len(curr)
				ret = ret[:0]
			}
			tmp := []string{beginWord}
			for _, idx := range curr {
				tmp = append(tmp, wordList[idx])
			}
			ret = append(ret, tmp)
			return
		}
		if len(curr) >= maxSize {
			return
		}
		for _, next := range nextGroup[start] {
			if seen[next] {
				continue
			}
			seen[next] = true
			helper(next, append(curr, next))
			seen[next] = false
		}
	}
	for i, w := range wordList {
		if !similar(beginWord, w) {
			continue
		}
		seen[i] = true
		helper(i, []int{i})
		seen[i] = false
	}
	return ret
}
