package main

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	count := func(v int) int { // smaller than v
		var ret int
		for i := 1; i <= m; i++ {
			ret += min(v/i, n)
		}
		return ret
	}
	for l < r {
		mid := (l + r) / 2
		if count(mid) < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
