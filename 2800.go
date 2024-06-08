package main

import (
	"strings"
)

func minimumString(a string, b string, c string) string {
	var ret string
	find := func(left, right string) int { // a last x == b first x
		dict := map[string]bool{}
		for i := 0; i <= len(right); i++ {
			dict[right[:i]] = true
		}
		for j := 0; j <= len(left); j++ {
			if dict[left[j:]] {
				return len(left[j:])
			}
		}
		return 0
	}
	in := func(a, b string) bool {
		if len(a) < len(b) {
			return false
		}
		for start, i := 0, len(b); i < len(a); start, i = start+1, i+1 {
			if a[start:i] == b {
				return true
			}
		}
		return false
	}
	merge := func(left, right string) string {
		if in(left, right) {
			return left
		}
		size := find(left, right)
		return left + right[size:]
	}
	helper := func(mid, left, right string) {
		sum := merge(merge(left, mid), right)
		if ret == "" || len(sum) < len(ret) || (len(sum) == len(ret) && strings.Compare(sum, ret) < 0) {
			ret = sum
		}
	}
	helper(a, b, c)
	helper(a, c, b)
	helper(b, a, c)
	helper(b, c, a)
	helper(c, a, b)
	helper(c, b, a)

	return ret
}
