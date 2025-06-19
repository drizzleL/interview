package main

import "math"

func minTravelTime(l int, n int, k int, position []int, time []int) int {
	cache := map[[4]int]int{}
	var helper func(i int, speed int, extraSpeed int, ops int) int
	helper = func(i int, speed int, extraSpeed int, ops int) (ret int) {
		if i == n-1 {
			if ops == 0 {
				return 0
			}
			return math.MaxInt32
		}
		if c, ok := cache[[4]int{i, speed, extraSpeed, ops}]; ok {
			return c
		}
		defer func() {
			cache[[4]int{i, speed, extraSpeed, ops}] = ret
		}()
		ret = math.MaxInt32
		nextPos := position[i+1] - position[i]
		ret = min(ret, (extraSpeed+time[i])*nextPos+helper(i+1, extraSpeed+time[i], 0, ops))
		if i != 0 && ops != 0 {
			ret = min(ret, speed*nextPos+helper(i+1, speed, extraSpeed+time[i], ops-1))
		}
		return
	}
	return helper(0, 0, 0, k)
}
