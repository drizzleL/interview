package main

func longestPalindrome3(s string, t string) int {
	check := func(x string) bool {
		for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
			if x[i] != x[j] {
				return false
			}
		}
		return true
	}
	var ret int
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			for k := 0; k < len(t); k++ {
				for l := k; l <= len(t); l++ {
					if check(s[i:j] + t[k:l]) {
						ret = max(ret, j-i+l-k)
					}
				}
			}
		}
	}
	return ret
}
