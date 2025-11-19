package main

func minLengthAfterRemovals2(s string) int {
	var cnt [2]int
	for _, c := range s {
		cnt[c-'a'] += 1
	}
	return len(s) - min(cnt[0], cnt[1])*2
}
