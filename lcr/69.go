package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		}
		if mid == 0 {
			l = mid + 1
			continue
		}
		if arr[mid] > arr[mid-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
