package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	dict := make([]int, 26)
	for _, c := range s {
		dict[c-'a']++
	}
	var cnt int
	for _, v := range dict {
		if v%2 == 1 {
			cnt++
		}
	}
	return cnt <= k
}
