package main

func maxValue(n int, index int, maxSum int) int {
	getSum := func(top int, size int) int {
		size2 := min(top-1, size)
		return (top-1+max(1, top-size))*size2/2 + (size - size2)
	}
	check := func(x int) bool {
		sum := x + getSum(x, index) + getSum(x, n-index-1)
		return sum <= maxSum
	}
	l, r := 1, maxSum
	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid - 1
			continue
		}
		if l+1 == r {
			if check(r) {
				return r
			}
			return l
		}
		l = mid
	}
	return l
}
