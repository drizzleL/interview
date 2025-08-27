package main

func countMaxOrSubsets(nums []int) int {
	var or int
	for _, num := range nums {
		or |= num
	}
	var ret int
	for mask := 1; mask < 1<<len(nums); mask++ {
		var or2 int
		for i := 0; i < len(nums); i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			or2 |= nums[i]
		}
		if or2 == or {
			ret += 1
		}
	}
	return ret
}
