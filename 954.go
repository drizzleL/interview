package main

import "sort"

func canReorderDoubled(arr []int) bool {
	var pos, neg []int
	var zeroCnt int
	for _, num := range arr {
		switch {
		case num > 0:
			pos = append(pos, num)
		case num < 0:
			neg = append(neg, -num)
		default:
			zeroCnt += 1
		}
	}
	if zeroCnt%2 != 0 {
		return false
	}
	check := func(nums []int) bool {
		if len(nums)%2 != 0 {
			return false
		}
		var vals []int
		dict := map[int]int{}
		for _, num := range nums {
			if dict[num] == 0 {
				vals = append(vals, num)
			}
			dict[num] += 1
		}
		sort.Ints(vals)
		for _, val := range vals {
			if dict[val] == 0 { // used before
				continue
			}
			if dict[val*2] < dict[val] {
				return false
			}
			dict[val*2] -= dict[val]
		}
		return true
	}
	return check(pos) && check(neg)
}
