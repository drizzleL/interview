package main

import "sort"

func minAbsDifference(nums []int, goal int) int {
	ret := abs(goal)
	uniq := func(arr []int) []int {
		var i int
		for j := 0; j < len(arr); j++ {
			if j != 0 && arr[j] == arr[j-1] {
				continue
			}
			arr[i] = arr[j]
			i++
		}
		arr = arr[:i]
		return arr
	}
	getPoss := func(vals []int) []int {
		ret := make([]int, 1<<len(vals))
		for i := range vals {
			ret[1<<i] = vals[i]
		}
		for i := range ret {
			ret[i] = vals[i&-i] + vals[i&(i-1)]
		}
		sort.Ints(ret)
		ret = uniq(ret)
		return ret
	}
	arr1, arr2 := getPoss(nums[:len(nums)/2]), getPoss(nums[len(nums)/2:])
	helper := func(v int, arr []int) int {
		idx := sort.SearchInts(arr, v)
		if idx == len(arr) {
			return abs(v - arr[len(arr)-1])
		}
		ret := abs(v - arr[idx])
		if idx != 0 {
			ret = min(ret, abs(v-arr[idx-1]))
		}
		return ret
	}
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	for _, num1 := range arr1 {
		ret = min(ret, helper(goal-num1, arr2))
	}
	return ret
}
