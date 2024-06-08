package main

func minimumPerimeter(neededApples int64) int64 {
	getSum := func(r int) int {
		return r * (r + 1) * (2*r + 1) * 2
	}
	l, r := 0, 1000000
	for l < r {
		mid := (l + r) / 2
		sum := getSum(mid)
		if sum < int(neededApples) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return int64(l) * 8
}
