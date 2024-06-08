package main

import "sort"

func permutation2(S string) []string {
	s := []byte(S)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	vis := make([]bool, len(s))
	var ret []string
	var helper func(tmp []byte)
	helper = func(tmp []byte) {
		if len(tmp) == len(s) {
			ret = append(ret, string(tmp))
			return
		}
		for i := range s {
			if vis[i] {
				continue
			}
			if i != 0 && s[i-1] == s[i] && !vis[i-1] {
				continue
			}
			vis[i] = true
			helper(append(tmp, s[i]))
			vis[i] = false
		}
	}
	helper(nil)
	return ret
}
