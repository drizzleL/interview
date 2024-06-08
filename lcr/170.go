package main

func reversePairs(record []int) int {
	var ret int
	merge := func(l1, l2 []int) []int {
		var l []int
		var i, j int
		for i < len(l1) && j < len(l2) {
			if l1[i] > l2[j] {
				ret += len(l2) - j
				l = append(l, l1[i])
				i++
			} else {
				l = append(l, l2[j])
				j++
			}
		}
		for i < len(l1) {
			l = append(l, l1[i])
			i++
		}
		for j < len(l2) {
			l = append(l, l2[j])
			j++
		}
		return l
	}
	var sort func(l []int) []int
	sort = func(l []int) []int {
		if len(l) <= 1 {
			return l
		}
		mid := len(l) / 2
		l1, l2 := l[:mid], l[mid:]
		l1, l2 = sort(l1), sort(l2)
		return merge(l1, l2)
	}
	sort(record)
	return ret
}
