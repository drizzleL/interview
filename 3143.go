package main

func maxPointsInsideSquare(points [][]int, s string) int {
	var l, r int
	for _, p := range points {
		r = max(r, abs(p[0]))
		r = max(r, abs(p[1]))
	}
	check := func(x int) (bool, int) {
		var ret int
		dict := [26]bool{}
		for i, p := range points {
			if abs(p[0]) > x || abs(p[1]) > x {
				continue
			}
			if dict[s[i]-'a'] {
				return false, 0
			}
			dict[s[i]-'a'] = true
			ret += 1
		}
		return true, ret
	}
	for l < r {
		mid := (l + r + 1) / 2
		if ok, _ := check(mid); ok {
			l = mid
		} else {
			r = mid - 1
		}
	}
	_, ret := check(l)
	return ret
}
