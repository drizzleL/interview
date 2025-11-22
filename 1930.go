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

func countPalindromicSubsequence2(s string) int {
	var pre, after [26]int
	var seen [26 * 26]bool
	for _, c := range s {
		after[c-'a'] += 1
	}
	for _, c := range s {
		i := int(c - 'a')
		after[c-'a'] -= 1
		for j := 0; j < 26; j++ {
			if pre[j] > 0 && after[j] > 0 {
				seen[i*26+j] = true
			}
		}
		pre[c-'a'] += 1
	}
	var ret int
	for _, v := range seen {
		if v {
			ret += 1
		}
	}
	return ret
}
