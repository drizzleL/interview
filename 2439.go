package main

func minimizeArrayValue(nums []int) int {
	var l, r int
	for _, num := range nums {
		r = max(r, num)
	}
	check := func(x int) bool {
		var tmp int
		for i := len(nums) - 1; i >= 0; i-- {
			tmp += nums[i]
			if tmp <= x {
				tmp = 0
				continue
			}
			tmp = tmp - x
		}
		return tmp <= 0
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
