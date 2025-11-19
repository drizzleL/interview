package main

func goodTriplets(nums1 []int, nums2 []int) int64 {
	idxs := make([]int, len(nums2))
	for i, v := range nums2 {
		idxs[v] = i
	}
	pre, suff := make([]int, len(nums1)), make([]int, len(nums1))
	preIdxs, suffIdxs := NewBiTree(len(nums2)), NewBiTree(len(nums1))
	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		pre[i] = preIdxs.Query(idxs[num])
		preIdxs.Update(idxs[num], 1)
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		num := nums1[i]
		suff[i] = len(nums1) - 1 - i - suffIdxs.Query(idxs[num])
		suffIdxs.Update(idxs[num], 1)
	}
	var ret int
	for i := range nums1 {
		ret += pre[i] * suff[i]
	}
	return int64(ret)
}
