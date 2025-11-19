package main

func maxOperations98(s string) int {
	var cnt, ret int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt += 1
			continue
		}
		if i != 0 && s[i-1] == '0' {
			continue
		}
		ret += cnt
	}
	return ret
}
