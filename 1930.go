package main

func countPalindromicSubsequence(s string) int {
	var dict1, dict2 [26]int
	for k := range dict1 {
		dict1[k] = -1
	}
	sums := make([][26]int, len(s)+1)
	for i, c := range s {
		if dict1[c-'a'] == -1 {
			dict1[c-'a'] = i
		}
		dict2[c-'a'] = i
		sums[i+1] = sums[i]
		sums[i+1][c-'a'] += 1
	}
	var ret int
	for k := range dict1 {
		if dict1[k] == -1 {
			continue
		}
		a, b := dict1[k], dict2[k]
		if b-a <= 1 {
			continue
		}
		for i := 0; i < 26; i++ {
			if sums[b][i]-sums[a+1][i] == 0 {
				continue
			}
			ret += 1
		}
	}
	return ret
}
