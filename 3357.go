package main

import (
	"math"
)

func minDifference2(nums []int) int {
	minNum, maxNum := math.MaxInt32, 0
	var maxGap int
	for i := 1; i < len(nums); i++ {
		if min(nums[i], nums[i-1]) == -1 && max(nums[i], nums[i-1]) != -1 {
			minNum = min(minNum, max(nums[i], nums[i-1]))
			maxNum = max(maxNum, max(nums[i], nums[i-1]))
		} else {
			maxGap = max(maxGap, abs(nums[i]-nums[i-1]))
		}
	}
	match := func(x1, x2 int, d int) bool {
		return abs(x1-x2) <= d
	}
	check := func(x, y int, d int) bool {
		var cnt, pre int
		for i := 0; i < len(nums); i++ {
			if nums[i] == -1 {
				if pre != 0 && cnt == 0 {
					if !match(pre, x, d) && !match(pre, y, d) {
						return false
					}
				}
				cnt += 1
				continue
			}
			if cnt >= 1 {
				if pre == 0 {
					if !match(nums[i], x, d) && !match(nums[i], y, d) {
						return false
					}
				} else {
					xMatch := match(pre, x, d) && match(nums[i], x, d)
					yMatch := match(pre, y, d) && match(nums[i], y, d)
					if !xMatch && !yMatch {
						if cnt == 1 {
							return false
						}
						if y-x > d {
							return false
						}
						if !match(min(pre, nums[i]), x, d) || !match(max(pre, nums[i]), y, d) {
							return false
						}
					}
				}
			}
			pre = nums[i]
			cnt = 0
		}
		return true
	}
	l, r := maxGap, (maxNum-minNum+1)/2
	for l < r {
		d := (l + r) / 2
		if check(minNum+d, maxNum-d, d) {
			r = d
		} else {
			l = d + 1
		}
	}
	return l
}
