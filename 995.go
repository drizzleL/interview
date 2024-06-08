package main

func minKBitFlips(nums []int, k int) int {
	flag := 1
	var ret int
	for i, num := range nums {
		if i >= k && nums[i-k] == 2 { // reverse at i - k
			flag = 1 - flag
		}
		if flag == num {
			continue
		}
		if len(nums)-i < k {
			return -1
		}
		flag = 1 - flag
		nums[i] = 2
		ret += 1
	}
	return ret
}
