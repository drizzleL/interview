package main

func sortArray(nums []int) []int {
	median := func(i, j, k int) {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[j] > nums[k] {
			nums[j], nums[k] = nums[k], nums[j]
		}
		if nums[i] > nums[k] {
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	var helper func(i, j int)
	helper = func(i, j int) {
		if i >= j {
			return
		}
		median(i, j, (i+j)/2)
		pivot := nums[i]
		l1, l2 := i, i
		r := j
		for l2 < r {
			v := nums[l2+1]
			switch {
			case v == pivot:
				l2++
			case v < pivot:
				nums[l1], nums[l2+1] = nums[l2+1], nums[l1]
				l1++
				l2++
			case v > pivot:
				nums[l2+1], nums[r] = nums[r], nums[l2+1]
				r--
			}
		}
		helper(i, l1-1)
		helper(l2+1, j)
	}
	helper(0, len(nums)-1)
	return nums
}

func median(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		c = a
	}
	if b > c {
		b = c
	}
	return b
}
