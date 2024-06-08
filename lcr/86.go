package main

func partition(s string) [][]string {
	check := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	var ret [][]string
	var helper func(i int, curr []string)
	helper = func(i int, curr []string) {
		if i == len(s) {
			dest := make([]string, len(curr))
			copy(dest, curr)
			ret = append(ret, dest)
			return
		}
		for j := i + 1; j <= len(s); j++ {
			if check(s[i:j]) {
				helper(j, append(curr, s[i:j]))
			}
		}
	}
	helper(0, nil)
	return ret
}
