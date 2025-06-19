package main

func alternatingSubarray(nums []int) int {
	ret := 0
	up, down := 0, 0
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] == nums[i-1]+1:
			up, down = down+1, 0
		case nums[i] == nums[i-1]-1 && up > 0:
			up, down = 0, up+1
		default:
			up, down = 0, 0
		}
		ret = max(ret, up)
		ret = max(ret, down)
	}
	if ret == 0 {
		return -1
	}
	return ret + 1
}
