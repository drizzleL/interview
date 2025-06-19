package main

func minCapability(nums []int, k int) int {
	l, r := nums[0], nums[0]
	for _, num := range nums {
		l = min(l, num)
		r = max(r, num)
	}
	check := func(x int) bool {
		var cnt int
		var flag bool
		for _, num := range nums {
			if flag { // reset
				flag = false
				continue
			}
			if num > x {
				continue
			}
			flag = true
			cnt += 1
		}
		return cnt >= k
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
