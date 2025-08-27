package main

import "sort"

func maxTotal(value []int, limit []int) int64 {
	var ele [][2]int
	for i := range value {
		ele = append(ele, [2]int{value[i], limit[i]})
	}
	sort.Slice(ele, func(i, j int) bool {
		if ele[i][1] == ele[j][1] {
			return ele[i][0] > ele[j][0]
		}
		return ele[i][1] < ele[j][1]
	})
	var ret int
	for i, j := 0, 0; i < len(ele) && j < len(ele); {
		e := ele[i]
		ret += e[0]
		cnt := i - j + 1
		for j < len(ele) && ele[j][1] <= cnt {
			j += 1
		}
		i = max(i+1, j)
	}
	return int64(ret)
}
