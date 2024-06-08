package main

func subarraySum(nums []int, k int) int {
	presums := map[int]int{}
	presums[0] = 1
	var ret int
	var sum int
	for _, num := range nums {
		sum += num
		ret += presums[sum-k]
		presums[sum]++
	}
	return ret
}
