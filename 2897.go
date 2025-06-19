package main

func maxSum(nums []int, k int) int {
	dict := [32]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; nums[i] != 0 && j < 32; j++ {
			dict[j] += nums[i] & 1
			nums[i] >>= 1
		}
	}
	var ret int
	for i := 0; i < k; i++ {
		var num int
		for j := 31; j >= 0; j-- {
			if dict[j] > 0 {
				num += 1 << j
				dict[j] -= 1
			}
		}
		ret += num * num
		ret %= 1e9 + 7
	}
	return ret
}
