package main

func scoreOfString(s string) int {
	var ret int
	for i := 1; i < len(s); i++ {
		ret += abs(int(s[i] - s[i-1]))
	}
	return ret
}
