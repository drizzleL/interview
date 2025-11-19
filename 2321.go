package main

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	var sum1, sum2 int
	var presum, minPre, maxPre int
	var a, b int
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		presum += nums1[i] - nums2[i]
		a = max(a, presum-minPre)
		b = min(b, presum-maxPre)
		minPre = min(minPre, presum)
		maxPre = max(maxPre, presum)
	}
	return max(sum1+a, sum2+b)
}
