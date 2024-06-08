package main

func longestMountain(arr []int) int {
	var ret int
	var up, down int
	check := func() {
		if up > 0 && down > 0 {
			ret = max(ret, up+down+1)
		}
	}
	for i := range arr {
		if i == 0 {
			continue
		}
		if arr[i] == arr[i-1] {
			check()
			up = 0
			down = 0
			continue
		}
		if arr[i] > arr[i-1] {
			if down == 0 {
				up += 1
			} else {
				check()
				down = 0
				up = 1
			}
		} else {
			if up == 0 {
				down = 0
			} else {
				down += 1
				check()
			}
		}
	}
	return ret
}
