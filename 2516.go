package main

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	dict := make([][]int, 3)
	rightDict := make([]int, 3)
	for i := len(s) - 1; i >= 0; i-- {
		rightDict[s[i]-'a'] += 1
	}
	for _, v := range rightDict {
		if v < k {
			return -1
		}
	}
	ret := len(s)
	for i := 0; i < len(s); i++ {
		rightDict[s[i]-'a'] -= 1
		dict[s[i]-'a'] = append(dict[s[i]-'a'], i)
		idx := -1
		for j := 0; j <= 2; j++ {
			if rightDict[j] >= k {
				continue
			}
			idx = max(idx, dict[j][k-rightDict[j]-1])
		}
		ret = min(ret, len(s)-(i-idx))
	}
	return ret
}
