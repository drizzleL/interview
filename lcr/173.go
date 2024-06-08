package main

func takeAttendance(records []int) int {
	// records[i] = i
	i, j := 0, len(records)-1
	for i < j {
		mid := (i + j) / 2
		if records[mid] == mid {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}
