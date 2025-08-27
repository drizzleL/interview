package main

func maxValue5(nums []int) []int {
	dp := make([]int, len(nums))
	type ele struct {
		start          int
		maxVal, minVal int
	}
	var eles []*ele
	for i := len(nums) - 1; i >= 0; i-- {
		maxVal := nums[i]
		minVal := nums[i]
		for len(eles) != 0 {
			top := eles[len(eles)-1]
			if maxVal >= top.maxVal || maxVal > top.minVal {
				maxVal = max(maxVal, top.maxVal)
				minVal = min(minVal, top.minVal)
				eles = eles[:len(eles)-1]
				continue
			}
			break
		}
		eles = append(eles, &ele{i, maxVal, minVal})
	}
	i := len(nums) - 1
	for _, ele := range eles {
		for i >= ele.start {
			dp[i] = ele.maxVal
			i--
		}
	}
	return dp
}
