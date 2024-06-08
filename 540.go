package main

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] != nums[m-1] && nums[m] != nums[m+1] {
			return nums[m]
		}
		if nums[m] == nums[m-1] {
			if (m-l+1)%2 == 0 {
				l = m + 1
			} else {
				r = m - 2
			}
		} else {
			if (r-m+1)%2 == 0 {
				r = m - 1
			} else {
				l = m + 2
			}
		}
	}
	return nums[l]
}
