package main

func minimumOperations2(nums []int) int {
	dict := map[int]bool{}
	idx := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if dict[nums[i]] {
			idx = i
			break
		}
		dict[nums[i]] = true
	}
	if idx == -1 {
		return 0
	}
	return idx/3 + 1
}
