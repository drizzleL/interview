package main

func subsetXORSum(nums []int) int {
	helper := func(mask int) int {
		var ret int
		for i := 0; i < len(nums); i++ {
			if (1<<i)&mask == 0 {
				continue
			}
			ret ^= nums[i]
		}
		return ret
	}
	var ret int
	for i := 1; i < 1<<len(nums); i++ {
		ret += helper(i)
	}
	return ret
}
