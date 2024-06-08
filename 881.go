package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)
	var ret int
	for l <= r {
		ret += 1
		if l == r {
			break
		}
		if people[l]+people[r] <= limit {
			l++
			r--
			continue
		} else {
			r--
		}
	}
	return ret
}
