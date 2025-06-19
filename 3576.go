package main

import "math"

func canMakeEqual(nums []int, k int) bool {
	helper := func(v int) int {
		flag := 1
		var ret int
		for _, num := range nums {
			num *= flag
			if num == v {
				flag = 1
				continue
			}
			ret += 1
			flag = -1
		}
		if flag == -1 {
			return math.MaxInt32
		}
		return ret
	}
	return min(helper(1), helper(-1)) <= k
}
