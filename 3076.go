package main

import (
	"sort"
	"strings"
)

func shortestSubstrings(arr []string) []string {
	getStrs := func(str string) []string {
		dict := map[string]bool{}
		var ret []string
		for i := 0; i < len(str); i++ {
			for j := len(str) - 1; j >= i; j-- {
				sub := str[i : j+1]
				if dict[sub] {
					break
				}
				dict[sub] = true
				ret = append(ret, sub)
			}
		}
		sort.Slice(ret, func(i, j int) bool {
			if len(ret[i]) == len(ret[j]) {
				return strings.Compare(ret[i], ret[j]) < 0
			}
			return len(ret[i]) < len(ret[j])
		})
		return ret
	}
	dict := map[string]int{}
	ret := make([]string, len(arr))
	for _, str := range arr {
		subStrs := getStrs(str)
		for _, sub := range subStrs {
			dict[sub] += 1
		}
	}
	for i, str := range arr {
		subStrs := getStrs(str)
		for _, sub := range subStrs {
			if dict[sub] == 1 {
				ret[i] = sub
				break
			}
		}
	}
	return ret
}
