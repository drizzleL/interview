package main

func oneEditAway(first string, second string) bool {
	find := func(a, b string) int {
		var i int
		for ; i < len(a); i++ {
			if a[i] != b[i] {
				break
			}
		}
		return i
	}
	if len(first) == len(second) {
		diffI := find(first, second)
		if diffI == len(first) { // same string
			return true
		}
		return first[diffI+1:] == second[diffI+1:]
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	diffI := find(first, second)
	return first[diffI:] == second[diffI+1:]
}
