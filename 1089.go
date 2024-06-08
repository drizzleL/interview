package main

func duplicateZeros(arr []int) {
	var cnt int
	var last, c int
	j := len(arr) - 1
	for last, c = range arr {
		if c == 0 {
			cnt += 2
		} else {
			cnt += 1
		}
		if cnt >= len(arr) {
			if cnt > len(arr) {
				last -= 1
				arr[len(arr)-1] = 0
				j -= 1
			}
			break
		}
	}
	for ; j >= 0; j-- {
		if arr[last] == 0 {
			arr[j] = 0
			arr[j-1] = 0
			j -= 1
		} else {
			arr[j] = arr[last]
		}
		last -= 1
	}
}
