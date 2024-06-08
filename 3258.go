package main

func countKConstraintSubstrings(s string, k int) int {
	if k == 0 {
		return 0
	}
	var cnt [2]int
	var i int
	var ret int
	for j := 0; j < len(s); j++ {
		cnt[s[j]-'0'] += 1
		for cnt[0] > k && cnt[1] > k {
			cnt[s[i]-'0'] -= 1
			i += 1
		}
		ret += j - i + 1
	}
	return ret
}
