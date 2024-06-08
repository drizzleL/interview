package main

import (
	"strconv"
	"strings"
)

func countPairs3(nums []int) int {
	var strs []string
	var maxLen int
	for _, num := range nums {
		str := strconv.Itoa(num)
		maxLen = max(maxLen, len(str))
		strs = append(strs, str)
	}
	for i, str := range strs {
		str = strings.Repeat("0", maxLen-len(str)) + str
		strs[i] = str
	}
	var check func(str []byte, dict map[string]bool, flag int)
	check = func(str []byte, dict map[string]bool, flag int) {
		dict[string(str)] = true
		if flag == 0 {
			return
		}
		for i := 0; i < len(str); i++ {
			for j := i + 1; j < len(str); j++ {
				str[i], str[j] = str[j], str[i]
				check(str, dict, flag-1)
				str[i], str[j] = str[j], str[i]
			}
		}
	}
	findAdj := func(str string) []string {
		dict := map[string]bool{}
		check([]byte(str), dict, 2)
		var ret []string
		for k := range dict {
			ret = append(ret, k)
		}
		return ret
	}
	var ret int
	dict := map[string]int{}
	for i := 0; i < len(strs); i++ {
		ret += dict[strs[i]]
		for _, str := range findAdj(strs[i]) {
			dict[str] += 1
		}
	}
	return ret
}
