package main

func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	var found bool
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt += 1
			if cnt >= k*2 {
				return true
			}
		} else {
			if cnt >= k {
				if found {
					return true
				}
				found = true
			} else {
				found = false
			}
			cnt = 1
		}
	}
	return cnt >= k && found
}
