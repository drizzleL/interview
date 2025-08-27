package main

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(arr []int) string {
	seen := make([]bool, 4)
	var vals [][2]int
	add := func(val []int) {
		a, b := val[0]*10+val[1], val[2]*10+val[3]
		if a > 23 || b > 59 {
			return
		}
		vals = append(vals, [2]int{a, b})
	}
	var dfs func(val []int)
	dfs = func(val []int) {
		if len(val) == 4 {
			add(val)
			return
		}
		for i := range arr {
			if seen[i] {
				continue
			}
			seen[i] = true
			dfs(append(val, arr[i]))
			seen[i] = false
		}
	}
	dfs(nil)
	sort.Slice(vals, func(i, j int) bool {
		if vals[i][0] == vals[j][0] {
			return vals[i][1] > vals[j][1]
		}
		return vals[i][0] > vals[j][0]
	})
	if len(vals) == 0 {
		return ""
	}
	return fmt.Sprintf("%d:%d", vals[0][0], vals[0][1])
}
