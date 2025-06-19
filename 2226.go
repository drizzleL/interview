package main

func maximumCandies(candies []int, k int64) int {
	var l, r int
	for _, c := range candies {
		r = max(r, c)
	}
	check := func(x int) bool {
		var cnt int
		for _, c := range candies {
			cnt += c / x
		}
		return cnt >= int(k)
	}
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
