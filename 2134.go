package main

func minSwaps(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	ret := len(nums)
	var now int
	for i := 0; i < sum; i++ {
		now += nums[i]
	}
	ret = min(ret, sum-now)
	if ret == 0 {
		return ret
	}
	for l, r := 0, sum; l < len(nums); l, r = l+1, (r+1)%len(nums) {
		now += nums[r]
		now -= nums[l]
		ret = min(ret, sum-now)
	}
	return ret
}
