package main

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	getKey := func(x int) int {
		if x < 0 {
			return (x+1)/(t+1) - 1
		}
		return x / (t + 1)
	}
	dict := map[int]int{}
	check := func(k int, x int) bool {
		val, ok := dict[k]
		if !ok {
			return false
		}
		if math.Abs(float64(val-x)) <= float64(t) {
			return true
		}
		return false
	}
	for i, num := range nums {
		key := getKey(num)
		if check(key, num) || check(key-1, num) || check(key+1, num) {
			return true
		}
		dict[key] = num
		if i >= k {
			delete(dict, getKey(nums[i-k]))
		}
	}
	return false
}
