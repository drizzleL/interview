package main

func canMakePaliQueries(s string, queries [][]int) []bool {
	dp := make([][26]int, len(s)+1)
	for i, c := range s {
		for j := 0; j < 26; j++ {
			dp[i+1][j] = dp[i][j]
		}
		dp[i+1][c-'a']++
	}
	var ret []bool
	for _, q := range queries {
		s, e, k := q[0], q[1], q[2]
		var c int
		for j := 0; j < 26; j++ {
			cnt := dp[e+1][j] - dp[s][j]
			if cnt%2 == 0 {
				continue
			}
			c++
		}
		ret = append(ret, (c)/2 <= k)
	}
	return ret
}
