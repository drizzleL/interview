package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ret := [][]int{{nums[0]}}
	insert := func(nums []int, v int) [][]int {
		var ret [][]int
		for i := 0; i < len(nums); i++ {
			tmp := make([]int, len(nums)+1)
			tmp[i] = v
			for j := 0; j < i; j++ {
				tmp[j] = nums[j]
			}
			for j := i + 1; j < len(tmp); j++ {
				tmp[j] = nums[j-1]
			}
			ret = append(ret, tmp)
		}
		return ret
	}
	helper := func(v int) {
		size := len(ret)
		for i := 0; i < size; i++ {
			ret = append(ret, insert(ret[i], v)...)
			ret[i] = append(ret[i], v)
		}
	}
	for i := 1; i < len(nums); i++ {
		helper(nums[i])
	}
	return ret
}
