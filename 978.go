package main

func maxTurbulenceSize(arr []int) int {
	ret, inc, dec := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			inc, dec = 1, 1
		} else if arr[i] > arr[i-1] {
			inc, dec = dec+1, 1
		} else {
			inc, dec = 1, inc+1
		}
		ret = max(ret, inc)
		ret = max(ret, dec)
	}
	return ret
}
