package main

func findKthLargest(nums []int, k int) int {
	var helper func(l, r int, k int) int
	helper = func(l, r int, k int) int {
		midl, midr := l, l
		pivot := nums[l]
		j := r
		for midr < j {
			v := nums[midr+1]
			switch {
			case v == pivot:
				midr += 1
			case v < pivot:
				nums[midr+1], nums[j] = nums[j], nums[midr+1]
				j -= 1
			case v > pivot:
				nums[midl], nums[midr+1] = nums[midr+1], nums[midl]
				midl += 1
				midr += 1
			}
		}
		if midl-l >= k {
			return helper(l, midl-1, k)
		}
		if midr-l+1 < k {
			return helper(midr+1, r, k-(midr-l+1))
		}
		return pivot
	}
	return helper(0, len(nums)-1, k)
}
