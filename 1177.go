package main

func canMakePaliQueries(s string, queries [][]int) []bool {
	dp := make([][26]int, len(s)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1]
		dp[i][s[i-1]-'a'] += 1
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		a, b := dp[q[1]+1], dp[q[0]]
		for j := 0; j < 26; j++ {
			a[j] -= b[j]
		}
		var cnt int
		for _, v := range a {
			if v%2 == 0 {
				continue
			}
			cnt += 1
		}
		if (q[1]-q[0]+1)%2 == 1 {
			cnt -= 1
		}
		cnt -= q[2] * 2
		ret[i] = cnt <= 0
	}
	return ret
}
