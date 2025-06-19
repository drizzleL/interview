package main

func maxIncreasingSubarrays(nums []int) int {
	ret := 1
	cnt := 1
	var lastCnt int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt += 1
			ret = max(ret, cnt/2)
			ret = max(ret, min(cnt, lastCnt))
		} else {
			lastCnt = cnt
			cnt = 1
		}
	}
	return ret
}
