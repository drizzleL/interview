package main

func longestOnes(nums []int, k int) int {
	var l int
	var ret int
	for i, num := range nums {
		if num == 0 {
			k--
		}
		if k < 0 {
			for nums[l] == 1 {
				l++
			}
			l++
			k++
		}
		ret = max(ret, i-l+1)
	}
	return ret
}
