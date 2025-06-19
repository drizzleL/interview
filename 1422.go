package main

func maxScore4(s string) int {
	var zeros, ones int
	for i := 0; i < len(s); i++ {
		ones += int(s[i] - '0')
	}
	var ret int
	for i := 0; i < len(s); i++ {
		zeros += 1 - int(s[i]-'0')
		ones -= int(s[i] - '0')
		ret = max(ret, zeros+ones)
	}
	return ret
}
