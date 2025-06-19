package main

import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	l, r := 0, len(tasks)
	sort.Ints(tasks)
	sort.Ints(workers)
	check := func(x int) bool {
		if x > len(workers) { // task > worker
			return false
		}
		a, b := tasks[:x], workers[len(workers)-x:]
		var extra []int
		p := pills
		for i, j := len(a)-1, len(b)-1; i >= 0; i-- {
			for j >= 0 && b[j]+strength >= a[i] {
				extra = append(extra, b[j])
				j--
			}
			if len(extra) == 0 {
				return false
			}
			if extra[0] >= a[i] {
				extra = extra[1:]
				continue
			}
			extra = extra[:len(extra)-1]
			p -= 1
			if p < 0 {
				return false
			}
		}
		return true
	}
	for l < r {
		mid := (l + r + 1) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
