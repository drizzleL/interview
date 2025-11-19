package main

func numSub(s string) int {
	var ret int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		j := i + 1
		for j < len(s) && s[j] == '1' {
			j += 1
		}
		ret += (j - i) * (j - i + 1) / 2
		i = j - 1
	}
	return ret
}
