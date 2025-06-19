package main

func partitionLabels(s string) []int {
	var ret []int
	last := make([]int, 26)
	for i, c := range s {
		last[c-'a'] = i
	}
	var segStart, segLast int
	for i, c := range s {
		currLast := last[c-'a']
		segLast = max(segLast, currLast)
		if segLast == i {
			ret = append(ret, segLast-segStart+1)
			segStart = i + 1
			segLast = i + 1
			continue
		}
	}
	return ret
}
