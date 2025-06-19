package main

func canSortArray(nums []int) bool {
	getBits := func(num int) int {
		var ret int
		for num != 0 {
			if num&1 != 0 {
				ret += 1
			}
			num >>= 1
		}
		return ret
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				continue
			}
			if getBits(nums[i]) != getBits(nums[j]) {
				return false
			}
		}
	}
	return true
}
