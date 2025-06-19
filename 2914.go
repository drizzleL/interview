package main

func minChanges(s string) int {
	var ret int
	for i := 0; i < len(s)/2; i++ {
		if s[i*2] != s[i*2+1] {
			ret += 1
		}
	}
	return ret
}
