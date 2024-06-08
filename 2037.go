package main

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var ret int
	for i := range seats {
		ret += abs(students[i] - seats[i])
	}
	return ret
}
