package main

func specialTriplets(nums []int) int {
	before, after := map[int]int{}, map[int]int{}
	for i := 2; i < len(nums); i++ {
		after[nums[i]]++
	}
	before[nums[0]] += 1
	var ret int
	for i := 1; i < len(nums)-1; i++ {
		ret += before[nums[i]*2] * after[nums[i]*2]
		ret %= 1e9 + 7
		before[nums[i]] += 1
		after[nums[i+1]] -= 1
	}
	return ret
}
