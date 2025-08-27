package main

func maxScore10(points []int, m int) int64 {
	check := func(x int) bool {
		var ret, last int
		for i := 0; i < len(points) && ret <= m; i++ {
			p := points[i]
			bar := (x + p - 1) / p
			if last >= bar {
				last = 0
				if i != len(points)-1 { // add to reach next
					ret += 1
				}
				continue
			}
			ret += (bar-last)*2 - 1
			last = bar - last - 1
		}
		return ret <= m
	}
	var l, r int
	for _, p := range points {
		r = max(r, p*m)
	}
	for l < r {
		mid := l + (r-l+1)/2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return int64(l)
}
