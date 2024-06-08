package main

func search(arr []int, target int) int {
	if len(arr) != 0 && arr[0] == target {
		return 0
	}
	var start int
	if len(arr) != 0 {
		v := arr[0]
		for len(arr) != 0 && arr[len(arr)-1] == v {
			arr = arr[:len(arr)-1]
		}
		for len(arr) != 0 && arr[0] == v {
			arr = arr[1:]
			start += 1
		}
	}
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] == target {
			r = mid
			continue
		}
		if arr[mid] > arr[r] { // mid in left bigger
			if arr[mid] > target && arr[0] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if arr[mid] < target && arr[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	if arr[l] == target {
		return l + start
	}
	return -1
}
