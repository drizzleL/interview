package main

import "strings"

func smallestNumber2(num string, t int64) string {
	if t == 1 {
		ret := strings.Repeat("1", len(num))
		for i := range ret {
			if num[i] > ret[i] {
				ret += "1"
				break
			}
		}
		return ret
	}
	primes := []int{2, 3, 5, 7}
	dict := []int{0, 0, 0, 0}
	m := int(t)
	var cnt int
	for i := 0; i < len(primes); i++ {
		for m%primes[i] == 0 {
			m /= primes[i]
			dict[i] += 1
			cnt += 1
		}
	}
	if m != 1 {
		return "-1"
	}
	if len(s) < cnt {
		return "23335577"
	}
	return ""
}
