package main

import "math"

func numSquarefulPerms(nums []int) int {
	check := func(a int) bool {
		sq := int(math.Sqrt(float64(a)))
		return sq*sq == a
	}
	end := 1<<len(nums) - 1
	var helper func(mask int, pre int) int
	helper = func(mask int, pre int) (ret int) {
		if mask == end { // reach end
			return 1
		}
		used := map[int]bool{}
		for j, num := range nums {
			if used[num] {
				continue
			}
			if mask&(1<<j) != 0 {
				continue
			}
			if pre != 0 && !check(pre+num) {
				continue
			}
			used[num] = true
			ret += helper(mask|(1<<j), num)
		}
		return
	}
	return helper(0, 0)
}
