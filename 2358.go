package main

func maximumGroups(grades []int) int {
	check := func(x int) bool {
		return (1+x)*x/2 <= len(grades)
	}
	l, r := 0, len(grades)
	for l < r {
		mid := (l + r + 1) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
