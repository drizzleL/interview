package main

import (
	"sort"
)

func maxSubstringLength(s string, k int) bool {
	var first, last, dict [26]int
	for i, c := range s {
		idx := c - 'a'
		if dict[idx] == 0 {
			first[idx] = i
		}
		last[idx] = i
		dict[idx] += 1
	}
	var spans [][2]int
	for a, i := range first {
		if dict[a] == 0 {
			continue
		}
		for b, j := range last {
			if dict[b] == 0 {
				continue
			}
			if i > j {
				continue
			}
			var cnt int
			for k, c := range dict {
				if c == 0 {
					continue
				}
				if first[k] >= i && last[k] <= j {
					cnt += c
				}
			}
			if cnt == j-i+1 && cnt != len(s) {
				spans = append(spans, [2]int{i, j})
			}
		}
	}
	sort.Slice(spans, func(i, j int) bool {
		return spans[i][1] < spans[j][1]
	})
	lastEnd := -1
	var cnt int
	for _, sp := range spans {
		if sp[0] > lastEnd {
			cnt += 1
			lastEnd = sp[1]
		}
	}
	return cnt >= k
}
