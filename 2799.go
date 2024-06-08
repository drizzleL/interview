package main

func countCompleteSubarrays(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	var hasNum int
	curr := map[int]int{}
	var ret int
	for l, r := 0, 0; l < len(nums); l++ {
		for ; r < len(nums) && hasNum < len(dict); r++ {
			curr[nums[r]] += 1
			if curr[nums[r]] == 1 {
				hasNum += 1
			}
		}
		if hasNum == len(dict) {
			ret += len(nums) - r + 1
		}
		curr[nums[l]] -= 1
		if curr[nums[l]] == 0 {
			hasNum -= 1
		}
	}
	return ret
}
