package main

func minSum(nums1 []int, nums2 []int) int64 {
	helper := func(nums []int) (sum, cnt int) {
		for _, num := range nums {
			sum += num
			if num == 0 {
				cnt += 1
			}
		}
		return
	}
	sum1, cnt1 := helper(nums1)
	sum2, cnt2 := helper(nums2)
	if cnt1 == 0 && sum1 < sum2+cnt2 {
		return -1
	}
	if cnt2 == 0 && sum2 < sum1+cnt1 {
		return -1
	}
	return int64(min(sum1+cnt1, sum2+cnt2))
}
