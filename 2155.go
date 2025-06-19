package main

func maxScoreIndices(nums []int) []int {
	var left, right int
	for _, num := range nums {
		if num == 1 {
			right += 1
		}
	}
	maxNum := left + right
	ret := []int{0}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left += 1
		} else {
			right -= 1
		}
		if left+right < maxNum {
			continue
		}
		if left+right > maxNum {
			ret = ret[:0]
		}
		maxNum = left + right
		ret = append(ret, i+1)
	}
	return ret
}
