package main

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	queryDict := map[int]int{}
	for i, q := range queries {
		queryDict[q] = i
	}
	var l, r int
	var servered int
	dict := make([]int, n+1)
	sortQ := make([]int, len(queries))
	copy(sortQ, queries)
	sort.Ints(sortQ)
	ansDict := map[int]int{}
	for _, q := range sortQ {
		for r < len(logs) {
			lg := logs[r]
			if lg[1] > q {
				break
			}
			if dict[lg[0]] == 0 {
				servered += 1
			}
			dict[lg[0]]++
			r++
		}
		for l < len(logs) {
			lg := logs[l]
			if lg[1] >= q-x {
				break
			}
			dict[lg[0]]--
			if dict[lg[0]] == 0 {
				servered--
			}
			l++
		}
		ansDict[q] = n - servered
	}
	var ret []int
	for _, q := range queries {
		ret = append(ret, ansDict[q])
	}
	return ret
}
