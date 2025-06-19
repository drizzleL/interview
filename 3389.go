package main

import "math"

func makeStringGood(s string) int {
	dict := [26]int{}
	for _, c := range s {
		dict[c-'a'] += 1
	}
	var maxFreq int
	for _, v := range dict {
		maxFreq = max(maxFreq, v)
	}
	check := func(dict [26]int, m int) int {
		ret := make([]int, 26)
		for i := 25; i >= 0; i-- {
			ret[i] = dict[i]
			if dict[i] > m {
				ret[i] = min(ret[i], dict[i]-m)
			} else {
				ret[i] = min(ret[i], m-dict[i])
			}
			if i != 25 {
				ret[i] += ret[i+1]
			}
			if i != 25 && dict[i+1] < m {
				var tmp int
				if i < 24 {
					tmp = ret[i+2]
				}
				diff1 := dict[i]
				if dict[i] > m {
					diff1 = dict[i] - m
				}
				diff := m - dict[i+1]
				x := diff1 + max(0, diff-diff1)
				ret[i] = min(ret[i], x+tmp)
			}
		}
		return ret[0]
	}
	ret := math.MaxInt32
	for i := 0; i <= maxFreq; i++ {
		ret = min(ret, check(dict, i))
	}
	return ret
}
