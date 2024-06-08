package main

func missingTwo(nums []int) []int {
	flag := 0b00
	var ret []int
	var helper func(i int)
	helper = func(i int) {
		val := nums[i]
		if val == i+1 { // already perfect
			return
		}
		if val == 0 {
			return
		}
		if val > len(nums) { // overflow, mark zero
			flag |= 1 << (val - len(nums) - 1)
			nums[i] = 0
			return
		}
		nums[val-1], nums[i] = nums[i], nums[val-1]
		helper(i)
	}
	for i := 0; i < len(nums); i++ {
		helper(i)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ret = append(ret, i+1)
		}
	}
	if flag&0b01 == 0 {
		ret = append(ret, len(nums)+1)
	}
	if flag&0b10 == 0 {
		ret = append(ret, len(nums)+2)
	}
	return ret
}
