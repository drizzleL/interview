package main

func countSubarrays5(nums []int, k int64) int64 {
	var ret, sum int
	for i, j := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		for sum*(i-j+1) >= int(k) {
			sum -= nums[i]
			i += 1
		}
		ret += j - i - 1
	}
	return int64(ret)
}
