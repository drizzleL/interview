package main

func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]string, len(str1))
	for i := range dp {
		dp[i] = make([]string, len(str2))
	}
	var helper func(i, j int) string
	helper = func(i, j int) (ret string) {
		if i == len(str1) && j == len(str2) {
			return ""
		}
		if i == len(str1) {
			return str2[j:]
		}
		if j == len(str2) {
			return str1[i:]
		}
		if dp[i][j] != "" {
			return dp[i][j]
		}
		defer func() {
			dp[i][j] = ret
		}()
		if str1[i] == str2[j] {
			return str1[i:i+1] + helper(i+1, j+1)
		}
		a, b := str1[i:i+1]+helper(i+1, j), str2[j:j+1]+helper(i, j+1)
		if len(a) < len(b) {
			return a
		}
		return b
	}
	return helper(0, 0)
}

func shortestCommonSupersequence2(str1 string, str2 string) string {
	dp := make([][][2]int, len(str1))
	for i := range dp {
		dp[i] = make([][2]int, len(str2))
	}
	var helper func(i, j int) [2]int
	helper = func(i, j int) (ret [2]int) {
		if i == len(str1) && j == len(str2) {
			return [2]int{0, 0}
		}
		if i == len(str1) {
			return [2]int{len(str2) - j, 2}
		}
		if j == len(str2) {
			return [2]int{len(str1) - i, 1}
		}
		if dp[i][j][0] != 0 {
			return dp[i][j]
		}
		defer func() {
			dp[i][j] = ret
		}()
		if str1[i] == str2[j] {
			tmp := helper(i+1, j+1)
			return [2]int{tmp[0] + 1, 3}
		}
		a, b := helper(i+1, j), helper(i, j+1)
		if a[0] <= b[0] {
			return [2]int{a[0] + 1, 1}
		}
		return [2]int{b[0] + 1, 2}
	}
	helper(0, 0)
	var b []byte
	var toStr func(i, j int)
	toStr = func(i, j int) {
		if i == len(str1) && j == len(str2) {
			return
		}
		if i == len(str1) {
			b = append(b, str2[j:]...)
			return
		}
		if j == len(str2) {
			b = append(b, str1[i:]...)
			return
		}
		switch dp[i][j][1] {
		case 3:
			b = append(b, str1[i])
			toStr(i+1, j+1)
		case 1:
			b = append(b, str1[i])
			toStr(i+1, j)
		case 2:
			b = append(b, str2[j])
			toStr(i, j+1)
		}
	}
	toStr(0, 0)
	return string(b)
}
