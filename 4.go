package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var kth func(k int, nums1, nums2 []int) int
	kth = func(k int, nums1, nums2 []int) int {
		if len(nums1) == 0 {
			return nums2[k]
		}
		if len(nums2) == 0 {
			return nums1[k]
		}
		if k == 0 {
			return min(nums1[0], nums2[0])
		}
		ia, ib := len(nums1)/2, len(nums2)/2
		va, vb := nums1[ia], nums2[ib]
		if ia+ib < k {
			if va < vb { // remove before ia
				return kth(k-ia-1, nums1[ia+1:], nums2)
			}
			return kth(k-ib-1, nums1, nums2[ib+1:])
		}
		if va > vb {
			return kth(k, nums1[:ia], nums2)
		}
		return kth(k, nums1, nums2[:ib])
	}
	size := len(nums1) + len(nums2)
	if size%2 == 1 {
		return float64(kth(size/2, nums1, nums2))
	}
	return (float64(kth(size/2-1, nums1, nums2)) + float64(kth(size/2, nums1, nums2))) / 2
}
