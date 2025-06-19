package main

import (
	"sort"
)

func shortestBeautifulSubstring(s string, k int) string {
	var ret []string
	var cnt int
	for i, j := 0, 0; j < len(s); i++ {
		for ; cnt < k && j < len(s); j++ {
			if s[j] == '0' {
				continue
			}
			cnt += 1
		}
		if cnt < k {
			break
		}
		for s[i] == '0' {
			i++
		}
		if len(ret) != 0 && j-i < len(ret[0]) {
			ret = ret[:0]
		}
		if len(ret) == 0 || len(ret[0]) == j-i {
			ret = append(ret, s[i:j])
		}
		cnt -= 1
	}
	if len(ret) == 0 {
		return ""
	}
	sort.Strings(ret)
	return ret[0]
}
