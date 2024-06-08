package main

func countTarget(scores []int, target int) int {
	find := func(v int) int {
		l, r := 0, len(scores)-1
		for l < r {
			mid := (l + r) / 2
			if scores[mid] < v {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	tIdx := find(target)
	if tIdx == len(scores) {
		return 0
	}
	rIdx := find(target + 1)
	return rIdx - tIdx
}
