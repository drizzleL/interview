package main

func getMaximumXor(nums []int, maximumBit int) []int {
	var v int
	for _, num := range nums {
		v ^= num
	}
	ret := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		k := 1<<maximumBit - 1
		ans := (k ^ v)
		ret = append(ret, ans)
		v ^= nums[i]
	}
	return ret
}
