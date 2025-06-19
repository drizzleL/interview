package main

func countAlternatingSubarrays(nums []int) int64 {
	var ret int
	for i := 0; i < len(nums); {
		j := i + 1
		for ; j < len(nums) && nums[j] != nums[j-1]; j++ {
		}
		size := j - i
		ret += (1 + size) * size / 2
		i = j
	}
	return int64(ret)
}
