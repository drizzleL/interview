package main

func subsets(nums []int) [][]int {
	var ret [][]int
	var helper func(i int, tmp []int)
	helper = func(i int, tmp []int) {
		if i == len(nums) {
			ret = append(ret, tmp)
			return
		}
		helper(i+1, tmp)
		helper(i+1, append(tmp, nums[i]))
	}
	helper(0, nil)
	return ret
}
