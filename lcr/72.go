package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x <= 3 {
		return 1
	}
	l, r := 2, x/2
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			r = mid - 1
			continue
		}
		if l == mid {
			if (mid+1)*(mid+1) == x {
				return mid + 1
			}
			return mid
		}
		l = mid
	}
	return l
}
