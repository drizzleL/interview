package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	if nums[l] == target {
		return true
	}
	for l < r && nums[l] == nums[r] {
		r--
	}
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
	}
	return false
}
