package main

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, 0
	for _, num := range nums {
		r = max(r, num)
	}
	check := func(x int) bool {
		var cnt int
		for _, num := range nums {
			cnt += num / x
			if num%x != 0 {
				cnt += 1
			}
		}
		return cnt <= threshold
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
