package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	once := make([]int, len(nums)+1)
	onceMax := make([][]int, len(nums)+1)
	for i := range onceMax {
		onceMax[i] = []int{0, 0}
	}
	var sum int
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		sum += nums[i]
	}
	for i := len(nums) - k; i >= 0; i-- {
		sum += nums[i]
		once[i] = sum
		onceMax[i] = onceMax[i+1]
		if once[i] >= onceMax[i][0] {
			onceMax[i] = []int{once[i], i}
		}
		sum -= nums[i+k-1]
	}
	twiceMax := make([][]int, len(nums))
	for i := range twiceMax {
		twiceMax[i] = []int{0, 0, 0}
	}
	for i := len(nums) - k - 1; i >= 0; i-- {
		twiceMax[i] = twiceMax[i+1]
		if once[i]+onceMax[i+k][0] >= twiceMax[i][0] {
			twiceMax[i] = []int{once[i] + onceMax[i+k][0], i, onceMax[i+k][1]}
		}
	}
	var retSum int
	var ret []int
	for i := 0; i < len(nums)-2*k; i++ {
		if once[i]+twiceMax[i+k][0] > retSum {
			retSum = once[i] + twiceMax[i+k][0]
			ret = []int{i, twiceMax[i+k][1], twiceMax[i+k][2]}
		}
	}
	return ret
}
