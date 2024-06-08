package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	var nums []int
	dict := map[int]int{}
	for _, h := range hand {
		if dict[h] == 0 {
			nums = append(nums, h)
		}
		dict[h] += 1
	}
	sort.Ints(nums)
	for _, num := range nums {
		for dict[num] != 0 {
			for i := 0; i < groupSize; i++ {
				if dict[num+i] == 0 {
					return false
				}
				dict[num+i] -= 1
			}
		}
	}
	return true
}
