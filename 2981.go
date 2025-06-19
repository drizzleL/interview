package main

import (
	"sort"
)

func maximumLength(s string) int {
	dict := [26][]int{}
	ret := -1
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		j := i
		for ; j+1 < len(s) && s[j] == s[j+1]; j++ {
		}
		dict[c] = append(dict[c], j-i+1)
		i = j
	}
	for i := range dict {
		sort.Ints(dict[i])
		if len(dict[i]) >= 3 {
			ret = max(ret, dict[i][len(dict[i])-3])
		}
		if len(dict[i]) >= 1 && dict[i][len(dict[i])-1] >= 3 {
			ret = max(ret, dict[i][len(dict[i])-1]-2)
		}
		if len(dict[i]) >= 2 && dict[i][len(dict[i])-1] >= 2 {
			ret = max(ret, min(dict[i][len(dict[i])-1]-1, dict[i][len(dict[i])-2]))
		}
	}
	return ret
}
