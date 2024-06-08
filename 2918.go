package main

func minSum(nums1 []int, nums2 []int) int64 {
	var sum1, sum2 int
	var zero1, zero2 int
	for _, num := range nums1 {
		sum1 += num
		if num == 0 {
			zero1++
		}
	}
	for _, num := range nums2 {
		sum2 += num
		if num == 0 {
			zero2++
		}
	}
	least1, least2 := sum1+zero1, sum2+zero2
	if zero1 > 0 && zero2 > 0 {
		return int64(max(least1, least2))
	}
	if zero1 == 0 && zero2 == 0 {
		if sum1 == sum2 {
			return int64(sum1)
		}
		return -1
	}
	if zero1 == 0 {
		if sum1 < least2 {
			return -1
		}
		return int64(sum1)
	}
	if zero2 == 0 {
		if sum2 < least1 {
			return -1
		}
		return int64(sum2)
	}
	return -1
}
