package main

import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	ret := make([]int, k)
	dict := map[int]int{}
	sort.Slice(logs, func(i, j int) bool {
		if logs[i][0] == logs[j][0] {
			return logs[i][1] < logs[j][1]
		}
		return logs[i][0] < logs[j][0]
	})
	for i, lg := range logs {
		if i != 0 && logs[i-1][0] == lg[0] && logs[i-1][1] == lg[1] {
			continue
		}
		dict[lg[0]] += 1
	}
	for _, v := range dict {
		if v == 0 {
			continue
		}
		ret[v-1] += 1
	}
	return ret
}
