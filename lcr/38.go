package main

func dailyTemperatures(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	var waitlist []int
	for i := 0; i < len(temperatures); i++ {
		for len(waitlist) != 0 && temperatures[waitlist[len(waitlist)-1]] < temperatures[i] {
			idx := waitlist[len(waitlist)-1]
			ret[idx] = i - idx
			waitlist = waitlist[:len(waitlist)-1]
		}
		waitlist = append(waitlist, i)
	}
	return ret
}
