package main

func maxFreqSum(s string) int {
	var dict [26]int
	var cnt1, cnt2 int
	for _, c := range s {
		dict[c-'a'] += 1
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			cnt1 = max(cnt1, dict[c-'a'])
		default:
			cnt2 = max(cnt2, dict[c-'a'])
		}
	}
	return cnt1 + cnt2
}
