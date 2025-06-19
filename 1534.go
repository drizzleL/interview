package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var ret int
	match := func(i, j, k int) bool {
		return abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c

	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if match(i, j, k) {
					ret += 1
				}
			}
		}
	}
	return ret
}
