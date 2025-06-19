package main

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i][0] < nums1[j][0]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i][0] < nums2[j][0]
	})
	var ret [][]int
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i == len(nums1) {
			ret = append(ret, nums2[j:]...)
			break
		}
		if j == len(nums2) {
			ret = append(ret, nums1[i:]...)
			break
		}
		switch {
		case nums1[i][0] == nums2[j][0]:
			ret = append(ret, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		case nums1[i][0] < nums2[j][0]:
			ret = append(ret, nums1[i])
			i++
		case nums1[i][0] > nums2[j][0]:
			ret = append(ret, nums2[j])
			j++
		}
	}
	return ret
}
