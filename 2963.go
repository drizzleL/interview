package main

func numberOfGoodPartitions(nums []int) int {
	dict := map[int]int{}
	for i, num := range nums {
		dict[num] = i
	}
	var maxRight int
	ret := 0
	for i, num := range nums {
		maxRight = max(maxRight, dict[num])
		if i == maxRight {
			if ret != 0 {
				ret *= 2
				ret %= 1e9 + 7
			} else {
				ret = 1
			}
			maxRight += 1
		}
	}
	return ret
}
