package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	helper := func(str string) string {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		return string(b)
	}
	dict := map[string][]string{}
	for _, str := range strs {
		dict[helper(str)] = append(dict[helper(str)], str)
	}
	var ret [][]string
	for _, ss := range dict {
		ret = append(ret, ss)
	}
	return ret
}
