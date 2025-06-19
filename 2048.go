package main

import (
	"strconv"
	"strings"
)

func nextBeautifulNumber(n int) int {
	base := []int{1, 22, 122, 333, 1333, 4444, 14444, 22333, 55555, 122333, 155555, 224444, 666666}
	ret := 1224444
	ns := strconv.Itoa(n)
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	nextPerm := func(x string) string {
		b := []byte(x)
		for i := len(x) - 2; i >= 0; i-- {
			if b[i] >= b[i+1] {
				continue
			}
			reverse(b[i+1:])
			for m := i + 1; m < len(b); m++ {
				if b[m] > b[i] {
					b[m], b[i] = b[i], b[m]
					return string(b)
				}
			}
		}
		return string(b)
	}
	for _, v := range base {
		s := strconv.Itoa(v)
		if len(s) < len(ns) {
			continue
		}
		if len(s) > len(ns) {
			ret = min(ret, v)
			continue
		}
		for {
			if strings.Compare(s, ns) > 0 {
				v2, _ := strconv.Atoi(s)
				ret = min(ret, v2)
			}
			next := nextPerm(s)
			if next == s {
				break
			}
			s = next
		}
	}
	return ret
}
