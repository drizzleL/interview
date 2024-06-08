package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ret []string
	check := func(s string) bool {
		num, _ := strconv.Atoi(s)
		if num > 255 {
			return false
		}
		return s == strconv.Itoa(num)
	}
	var helper func(i int, curr []string)
	helper = func(i int, curr []string) {
		if i == len(s) && len(curr) == 4 {
			ret = append(ret, strings.Join(curr, "."))
			return
		}
		if len(curr) == 4 {
			return
		}
		for j := i + 1; j <= len(s) && j <= i+3; j++ {
			if check(s[i:j]) {
				helper(j, append(curr, s[i:j]))
			}
		}
	}
	helper(0, nil)
	return ret
}
