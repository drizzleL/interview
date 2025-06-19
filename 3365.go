package main

func isPossibleToRearrange(s string, t string, k int) bool {
	size := len(s) / k
	dict := map[string]int{}
	for i := 0; i < len(s); i += size {
		j := i + size
		dict[s[i:j]] += 1
	}
	for i := 0; i < len(s); i += size {
		j := i + size
		if dict[t[i:j]] == 0 {
			return false
		}
		dict[t[i:j]] -= 1
	}
	return true
}
