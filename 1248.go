package main

func numberOfSubarrays(nums []int, k int) int {
	var odds []int
	for i, num := range nums {
		if num%2 == 1 {
			odds = append(odds, i)
		}
	}
	var ret int
	for i := k - 1; i < len(odds); i++ {
		var left, right int
		if i-k-1 == 0 {
			left = odds[i-k-1] + 1
		} else {
			left = odds[i-k-1] - odds[i-k-2]
		}
		if i == len(odds)-1 {
			right = len(nums) - odds[i]
		} else {
			right = odds[len(odds)-1] - odds[i]
		}
		ret += left * right
	}
	return ret
}
