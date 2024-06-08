package main

func countStudents(students []int, sandwiches []int) int {
	var cnt [2]int
	for _, s := range students {
		cnt[s] += 1
	}
	var i int
	for i < len(students) {
		if cnt[sandwiches[i]] == 0 {
			break
		}
		cnt[sandwiches[i]] -= 1
		i++
	}
	return len(students) - i
}
