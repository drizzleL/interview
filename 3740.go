package main

import "math"

func minimumDistance(nums []int) int {
	dict := map[int][]int{}
	ret := math.MaxInt32
	for i, num := range nums {
		dict[num] = append(dict[num], i)
		if len(dict[num]) < 3 {
			continue
		}
		arr := dict[num]
		a, b := arr[len(arr)-3], arr[len(arr)-2]
		ret = min(ret, (b-a)+(i-b)+(i-a))
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
