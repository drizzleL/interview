package main

func kLengthApart(nums []int, k int) bool {
	last := -k - 1
	for i, num := range nums {
		if num == 0 {
			continue
		}
		if i-last <= k {
			return false
		}
		last = i
	}
	return true
}
