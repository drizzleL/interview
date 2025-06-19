package main

func zeroFilledSubarray(nums []int) int64 {
	var pre, ret int
	for i, num := range nums {
		if num != 0 {
			pre = i + 1
			continue
		}
		ret += i - pre + 1
	}
	return int64(ret)
}
