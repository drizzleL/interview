package main

func minCost(colors string, neededTime []int) int {
	var lastColor rune
	var ret int
	var subSum int
	var subMax int
	for i, c := range colors {
		if c == lastColor {
			subSum += neededTime[i]
			subMax = max(subMax, neededTime[i])
		}
		if c != lastColor {
			ret += subSum - subMax
			subSum = neededTime[i]
			subMax = neededTime[i]
		}
		lastColor = c
	}
	ret += subSum - subMax
	return ret
}
