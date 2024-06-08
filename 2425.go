package main

func xorAllNums(nums1 []int, nums2 []int) int {
	var ret int
	if len(nums1)%2 == 1 {
		for _, num := range nums2 {
			ret ^= num
		}
	}
	if len(nums2)%2 == 1 {
		for _, num := range nums1 {
			ret ^= num
		}
	}
	return ret
}
