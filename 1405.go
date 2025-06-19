package main

import "strings"

func longestDiverseString(a int, b int, c int) string {
	var helper func(a, b, c int, aa, bb, cc string) string
	helper = func(a, b, c int, aa, bb, cc string) string {
		if a < b {
			return helper(b, a, c, bb, aa, cc)
		}
		if b < c {
			return helper(a, c, b, aa, cc, bb)
		}
		if b == 0 {
			return strings.Repeat(aa, min(a, 2))
		}
		usea := min(a, 2)
		useb := 0
		if a-usea >= b {
			useb = 1
		}
		return strings.Repeat(aa, usea) + strings.Repeat(bb, useb) + helper(a-usea, b-useb, c, aa, bb, cc)
	}
	return helper(a, b, c, "a", "b", "c")
}
