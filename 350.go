package main

func intersect(nums1 []int, nums2 []int) []int {
	dict := make([]int, 1001)
	for _, num := range nums1 {
		dict[num] += 1
	}
	var ret []int
	for _, num := range nums2 {
		if dict[num] == 0 {
			continue
		}
		ret = append(ret, num)
		dict[num] -= 1
	}
	return ret
}
