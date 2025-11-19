package main

import (
	"strconv"
)

func nextBeautifulNumber(n int) int {
	base := []int{1, 22, 122, 333, 1333, 4444, 14444, 22333, 55555, 122333, 155555, 224444, 666666}
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	var bb [][]byte
	for _, v := range base {
		b := []byte(strconv.Itoa(v))
		reverse(b)
		bb = append(bb, b)
	}
	cmp := func(a, b []byte) int {
		if len(a) > len(b) {
			return 1
		}
		if len(a) < len(b) {
			return -1
		}
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				return 1
			}
			if a[i] < b[i] {
				return -1
			}
		}
		return 0

	}
	ret := 1224444
	ns := []byte(strconv.Itoa(n))
	nextPerm := func(b []byte) []byte {
		for i := len(b) - 2; i >= 0; i-- {
			if b[i] >= b[i+1] {
				continue
			}
			reverse(b[i+1:])
			for m := i + 1; m < len(b); m++ {
				if b[m] > b[i] {
					b[m], b[i] = b[i], b[m]
					return b
				}
			}
		}
		return b
	}

	for _, v := range bb {
		if len(v) > len(ns)+1 {
			break
		}
		if cmp(v, ns) <= 0 {
			continue
		}
		reverse(v)
		for {
			v = nextPerm(v)
			if cmp(v, ns) <= 0 {
				continue
			}
			num, _ := strconv.Atoi(string(v))
			ret = min(ret, num)
			break
		}
	}
	return ret
}
