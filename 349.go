package main

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]bool)
	for _, num := range nums1 {
		dict[num] = true
	}
	for _, num := range nums2 {
		if dict[num] {
			dict[num] = false
		}
	}
	var ret []int
	for k, v := range dict {
		if !v {
			ret = append(ret, k)
		}
	}
	return ret
}
