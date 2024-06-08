package main

func partition(s string) [][]string {
	isPali := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	dp := make([][][]string, len(s)+1)
	dp[0] = [][]string{[]string{}}
	for j := range s {
		for i := 0; i <= j; i++ {
			if !isPali(i, j) {
				continue
			}
			for _, v := range dp[i] {
				tmp := make([]string, len(v), len(v)+len(s[i:j+1]))
				copy(tmp, v)
				tmp = append(tmp, s[i:j+1])
				dp[j+1] = append(dp[j+1], tmp)
			}
		}
	}
	return dp[len(s)]
}
