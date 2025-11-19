package main

func minimumTime77(d []int, r []int) int64 {
	check := func(x int) bool {
		d1, d2 := d[0], d[1]
		a, b := x/r[0], x/r[1]
		g := x / (lcm(r[0], r[1]))
		both := x - (a + b - g)
		onlyA := x - a - both
		onlyB := x - b - both
		d1 = max(0, d1-onlyA)
		d2 = max(0, d2-onlyB)
		return both >= d1+d2
	}
	low, high := d[0]+d[1], (d[0]+d[1])*2
	for low < high {
		mid := (low + high) / 2
		if check(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return int64(low)
}
