package main

func numSub(s string) int {
	var i int
	var ret int
	for i < len(s) {
		if s[i] == '0' {
			i++
			continue
		}
		oldI := i
		for i < len(s) && s[i] == '1' {
			i++
		}
		size := i - oldI
		ret += (1 + size) * size / 2
	}
	return ret
}
