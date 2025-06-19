package main

import "sort"

func maxSubstrings(word string) int {
	dict := [26][]int{}
	for i, c := range word {
		dict[c-'a'] = append(dict[c-'a'], i)
	}
	dp := make([]int, len(word)+1)
	for i := 0; i < len(word); i++ {
		dp[i+1] = dp[i]
		c := int(word[i] - 'a')
		idx := sort.SearchInts(dict[c], i-2) - 1
		if idx >= 0 && dict[c][idx] != i {
			dp[i+1] = max(dp[i+1], dp[dict[c][idx]]+1)
		}
	}
	return dp[len(word)]
}
