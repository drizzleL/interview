package main

func distinctSubseqII(s string) int {
	var cnt [26]int
	var ret int
	for _, c := range s {
		idx := int(c - 'a')
		oldCnt := cnt[idx]
		cnt[idx] = ret + 1
		cnt[idx] %= 1e9 + 7
		ret += cnt[idx] - oldCnt
		if ret < 0 {
			ret += 1e9 + 7
		}
		ret %= 1e9 + 7
	}
	return ret
}
