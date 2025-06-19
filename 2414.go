package main

func longestContinuousSubstring(s string) int {
	ret := 1
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			cnt += 1
		} else {
			cnt = 1
		}
		ret = max(ret, cnt)
	}
	return ret
}
