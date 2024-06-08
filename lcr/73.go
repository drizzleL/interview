package main

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	for l < r {
		mid := (l + r) / 2
		var hour int
		for _, p := range piles {
			hour += p / mid
			if p%mid != 0 {
				hour++
			}
		}
		if hour <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
