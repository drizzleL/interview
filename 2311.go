package main

func longestSubsequence2(s string, k int) int {
	var ret, val int
	pow := 1
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch c {
		case '0':
			ret += 1
		case '1':
			if val+pow <= k {
				val += pow
				ret += 1
			}
		}
		if val+pow <= k {
			pow *= 2
		}
	}
	return ret
}
