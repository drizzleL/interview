package main

func checkSubarraySum(nums []int, k int) bool {
	dict := map[int]int{}
	dict[0] = -1
	var sum int
	for i, num := range nums {
		sum += num
		sum %= k
		oldIdx, ok := dict[sum]
		if !ok {
			dict[sum] = i
			continue
		}
		if i-oldIdx >= 2 {
			return true
		}
	}
	return false
}
