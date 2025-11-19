package main

func maxSlidingWindow(nums []int, k int) []int {
	var q []int
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
	}
	ret := make([]int, len(nums)-k+1)
	ret[0] = q[0]
	for i := k; i < len(nums); i++ {
		for len(q) > 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
		if q[0] == nums[i-k] {
			q = q[1:]
		}
		ret[i-k+1] = q[0]
	}
	return ret
}
