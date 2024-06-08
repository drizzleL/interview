package main

import "math"

func minSessions(tasks []int, sessionTime int) int {
	ret := math.MaxInt32
	var solve func(i int, sessions []int)
	solve = func(i int, sessions []int) {
		if i == len(tasks) {
			ret = min(ret, len(sessions))
			return
		}
		if len(sessions) >= ret {
			return
		}
		for j, session := range sessions {
			if session+tasks[i] > sessionTime {
				continue
			}
			sessions[j] += tasks[i]
			solve(i+1, sessions)
			sessions[j] -= tasks[i]
		}
		sessions = append(sessions, tasks[i])
		solve(i+1, sessions)
	}
	solve(0, nil)
	return ret
}
