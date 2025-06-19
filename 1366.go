package main

import "sort"

func rankTeams(votes []string) string {
	dict := make([][]int, 26)
	for i := range dict {
		dict[i] = make([]int, len(votes[0]))
	}
	for _, vote := range votes {
		for i, c := range vote {
			dict[c-'a'][i] += 1
		}
	}
	b := []byte(votes[0])
	sort.Slice(b, func(i, j int) bool {
		c1, c2 := b[i], b[j]
		for m := 0; m < len(votes[0]); m++ {
			if dict[c1-'a'][m] == dict[c2-'a'][m] {
				continue
			}
			return dict[c1-'a'][m] > dict[c2-'a'][m]
		}
		return c1 > c2
	})
	return string(b)
}
