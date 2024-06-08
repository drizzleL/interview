package main

func subsets(nums []int) [][]int {
	var ret [][]int
	var helper func(tmp []int, i int)
	helper = func(tmp []int, i int) {
		if i == len(nums) {
			dst := make([]int, len(tmp))
			copy(dst, tmp)
			ret = append(ret, dst)
			return
		}
		helper(tmp, i+1)
		helper(append(tmp, nums[i]), i+1)
	}
	helper(nil, 0)
	return ret
}
