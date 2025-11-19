package main

import "sort"

func sortVowels(s string) string {
	vowelDict := map[byte]bool{}
	for _, c := range "aeiouAEIOU" {
		vowelDict[byte(c)] = true
	}
	t := make([]byte, len(s))
	var tmp []byte
	for i := 0; i < len(s); i++ {
		if !vowelDict[s[i]] {
			continue
		}
		tmp = append(tmp, s[i])
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	for i, j := 0, 0; i < len(s); i++ {
		if !vowelDict[s[i]] {
			t[i] = s[i]
			continue
		}
		t[i] = tmp[j]
		j++
	}
	return string(t)
}
