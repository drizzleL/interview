package main

func minOperations99(nums1 []int, nums2 []int) int {
	size := len(nums1)
	helper := func(v1, v2 int) (ret int) {
		for i := 0; i < size-1; i++ {
			if nums1[i] <= v1 && nums2[i] <= v2 {
				continue
			}
			if nums1[i] <= v2 && nums2[i] <= v1 {
				ret += 1
				continue
			}
			return -1
		}
		return ret
	}
	ret := helper(nums1[size-1], nums2[size-1])
	if ret == -1 {
		return -1
	}
	return min(ret, helper(nums2[size-1], nums1[size-1])+1)
}
