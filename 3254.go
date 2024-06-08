package main

func resultsArray(nums []int, k int) []int {
	ret := make([]int, len(nums)-k+1)
	var ascCnt int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			ascCnt += 1
		} else {
			ascCnt = 1
		}
		if i >= k-1 {
			if ascCnt >= k {
				ret[i-k+1] = nums[i]
			} else {
				ret[i-k+1] = -1
			}
		}
	}
	return ret
}
