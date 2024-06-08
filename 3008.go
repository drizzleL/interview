package main

import (
	"sort"
)

func beautifulIndices(s string, a string, b string, k int) []int {
	aMatches := getMatch(s, a)
	bMatches := getMatch(s, b)
	var ret []int
	for i := 0; i < len(aMatches); i++ {
		bl := sort.Search(len(bMatches), func(j int) bool {
			return bMatches[j] >= (aMatches[i] - k)
		})
		br := sort.Search(len(bMatches), func(j int) bool {
			return bMatches[j] > (aMatches[i] + k)
		})
		if br != bl {
			ret = append(ret, aMatches[i])
		}
	}
	return ret
}

func getMatch(s, t string) []int {
	next := getMatchNext(t)
	var ret []int
	for i, j := 0, 0; i < len(s); i++ {
		if t[j] == s[i] {
			j++
			if j == len(t) { // match all
				ret = append(ret, i-len(t)+1)
				j = next[len(t)-1]
			}
		} else {
			if j != 0 {
				i -= 1
				j = next[j-1]
			}
		}
	}
	return ret
}
func getMatchNext(t string) []int {
	next := make([]int, len(t))
	for i := 1; i < len(t); i++ {
		if t[i] == t[next[i-1]] {
			next[i] = next[i-1] + 1
		}
	}
	return next
}
