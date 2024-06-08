package main

func longestPalindrome(s string) int {
	dict := map[rune]int{}
	for _, c := range s {
		dict[c] += 1
	}
	var ret int
	var flag int
	for _, v := range dict {
		flag |= v % 2
		ret += v / 2 * 2
	}
	if flag == 1 {
		ret += 1
	}
	return ret
}
