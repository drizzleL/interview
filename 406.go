package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	ret := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		var j, cnt int
		for {
			if len(ret[j]) != 0 {
				j += 1
				continue
			}
			if cnt == people[i][1] {
				break
			}
			cnt += 1
			j += 1
		}
		ret[j] = people[i]
	}
	return ret
}
