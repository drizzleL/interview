package main

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, 0
	for _, num := range nums {
		r = max(r, num)
	}
	check := func(x int) bool {
		var m int
		for _, num := range nums {
			m += (num - 1) / x
		}
		return m <= maxOperations
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
