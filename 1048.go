package main

func longestStrChain(words []string) int {
	dict := map[int][]int{}
	var ret int
	dp := make([]int, len(words))
	for i, w := range words {
		dict[len(w)] = append(dict[len(w)], i)
	}
	check := func(a, b string) bool {
		var i, j int
		for ; i < len(a) && j < len(b); j++ {
			if a[i] == b[j] {
				i++
			}
		}
		return i == len(a)
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		}
		w := words[i]
		dp[i] = 1
		for _, j := range dict[len(w)+1] {
			if check(w, words[j]) {
				dp[i] = max(dp[i], dfs(j)+1)
			}
		}
		dp[i] += 1
		ret = max(ret, dp[i])
		return dp[i]
	}
	for i := 0; i < len(words); i++ {
		dfs(i)
	}
	return ret
}
