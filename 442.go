package main

func findDuplicates(nums []int) []int {
	var ret []int
	helper := func(i int) {
		for nums[i] > 0 && nums[i] != i+1 {
			v := nums[i]
			i2 := v - 1
			v2 := nums[i2]
			if v == v2 {
				ret = append(ret, v)
				nums[i2] *= -1
				break
			}
			nums[i], nums[i2] = nums[i2], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		helper(i)
	}
	return ret
}
