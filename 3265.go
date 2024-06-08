package main

import (
	"sort"
	"strconv"
	"strings"
)

func countPairs2(nums []int) int {
	sort.Ints(nums)
	var strs []string
	var maxLen int
	for _, num := range nums {
		str := strconv.Itoa(num)
		maxLen = max(maxLen, len(str))
		strs = append(strs, str)
	}
	var cnts [][10]int
	for i, str := range strs {
		str = strings.Repeat("0", maxLen-len(str)) + str
		strs[i] = str
		var cnt [10]int
		for _, c := range str {
			cnt[int(c-'0')] += 1
		}
		cnts = append(cnts, cnt)
	}
	cmp := func(i, j int) bool {
		if cnts[i] != cnts[j] {
			return false
		}
		var diff int
		for idx := range strs[i] {
			if strs[i][idx] != strs[j][idx] {
				diff += 1
			}
		}
		return diff <= 2
	}
	var ret int
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if cmp(i, j) {
				ret += 1
			}
		}
	}
	return ret
}
