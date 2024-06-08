package main

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		for nums[i] != 1 && i <= r {
			if nums[i] == 0 {
				nums[i], nums[l] = nums[l], nums[i]
				l++
				if i+1 == l {
					break
				}
			} else {
				nums[i], nums[r] = nums[r], nums[i]
				r--
			}
		}
	}
}
