package main

func numSplits(s string) int {
	var left, right [26]int
	var leftCnt, rightCnt int
	var ret int
	for _, c := range s {
		if right[c-'a'] == 0 {
			rightCnt += 1
		}
		right[c-'a'] += 1
	}
	for i := 0; leftCnt <= rightCnt; i++ {
		c := s[i]
		if left[c-'a'] == 0 {
			leftCnt += 1
		}
		left[c-'a'] += 1
		if right[c-'a'] == 1 {
			rightCnt -= 1
		}
		right[c-'a'] -= 1
		if leftCnt == rightCnt {
			ret += 1
		}
	}
	return ret
}
