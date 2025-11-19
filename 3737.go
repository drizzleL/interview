package main

func countMajoritySubarrays(nums []int, target int) int {
	var ret int
	isMajor := func(dict map[int]int, size int) bool {
		return dict[target] > size/2
	}
	for i := 0; i < len(nums); i++ {
		dict := map[int]int{}
		for j := i; j < len(nums); j++ {
			dict[nums[j]] += 1
			if isMajor(dict, j-i+1) {
				ret += 1
			}
		}
	}
	return ret
}
