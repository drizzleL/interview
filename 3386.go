package main

func buttonWithLongestTime(events [][]int) int {
	ret := events[0][0]
	takes := events[0][1]
	for i := 1; i < len(events); i++ {
		tmp := events[i][1] - events[i-1][1]
		if tmp > takes || (tmp == takes && events[i][0] < ret) {
			takes = tmp
			ret = events[i][0]
		}
	}
	return ret
}
