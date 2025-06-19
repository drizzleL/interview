package main

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	dp := make([][2]int, len(words))
	idx := len(words) - 1
	cmp := func(i, j int) bool {
		if groups[i] == groups[j] {
			return false
		}
		if len(words[i]) != len(words[j]) {
			return false
		}
		var dis int
		for k := range words[i] {
			if words[i][k] == words[j][k] {
				continue
			}
			dis += 1
		}
		return dis == 1
	}
	for i := len(words) - 1; i >= 0; i-- {
		dp[i] = [2]int{1, -1}
		for j := i + 1; j < len(words); j++ {
			if dp[j][0]+1 <= dp[i][0] {
				continue
			}
			if !cmp(i, j) {
				continue
			}
			dp[i][0] = dp[j][0] + 1
			dp[i][1] = j
		}
		if dp[i][0] > dp[idx][0] {
			idx = i
		}
	}
	var ret []string
	for idx != -1 {
		ret = append(ret, words[idx])
		idx = dp[idx][1]
	}
	return ret
}
