package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	findIdxs := func(str string) []int {
		var ret []int
		for i, c := range str {
			if c == '.' {
				ret = append(ret, i)
			}
		}
		return ret
	}
	dict := map[string]int{}
	for _, cpd := range cpdomains {
		seps := strings.Split(cpd, " ")
		cnt, _ := strconv.Atoi(seps[0])
		dict[seps[1]] += cnt
		for _, i := range findIdxs(seps[1]) {
			dict[seps[1][i+1:]] += cnt
		}

	}
	var ret []string
	for k, v := range dict {
		ret = append(ret, fmt.Sprintf("%d %s", v, k))
	}
	return ret
}
