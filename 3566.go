package main

func checkEqualPartitions(nums []int, target int64) bool {
	tt := int64(1)
	for _, num := range nums {
		tt *= int64(num)
	}
	if tt != target*target {
		return false
	}
	var ret bool
	var check func(i int, pre int64)
	check = func(i int, pre int64) {
		if pre == target {
			ret = true
			return
		}
		if i == len(nums) {
			return
		}
		check(i+1, pre*int64(nums[i]))
		check(i+1, pre)
	}
	check(0, 1)
	return ret
}
