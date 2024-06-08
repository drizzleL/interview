package main

func canTransform(start string, end string) bool {
	find := func(str string, b byte, i int) (next, cnt, start, end int) {
		start = i
		for ; i < len(str) && (str[i] == 'X' || str[i] == b); i++ {
			if str[i] == b {
				cnt += 1
				end = i
			}
		}
		next = i
		return
	}
	var i, j int
	for i < len(start) && start[i] == 'X' {
		i++
	}
	for j < len(end) && end[j] == 'X' {
		j++
	}
	var cnt1, cnt2, start1, start2, end1, end2 int
	for i < len(start) && j < len(end) {
		c := start[i]
		i, cnt1, start1, end1 = find(start, c, i)
		j, cnt2, start2, end2 = find(end, c, j)
		if cnt1 == 0 && cnt2 == 0 {
			return true
		}
		if cnt1 != cnt2 {
			return false
		}
		if c == 'L' && (start1 < start2 || end1 < end2) {
			return false
		}
		if c == 'R' && (start1 > start2 || end1 > end2) {
			return false
		}
	}
	return i == len(start) && j == len(end)
}
