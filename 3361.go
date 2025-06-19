package main

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	prePresum := make([]int, 27)
	for i := 0; i < len(previousCost); i++ {
		prePresum[i+1] = prePresum[i] + previousCost[i]
	}
	nextPresum := make([]int, 27)
	for i := 0; i < len(nextCost); i++ {
		nextPresum[i+1] = nextPresum[i] + nextCost[i]
	}
	var getNext func(i, j byte) int
	getNext = func(i, j byte) int {
		if i <= j {
			return nextPresum[j] - nextPresum[i]
		}
		return getNext(i, 26) + getNext(0, j) + nextCost[26]
	}
	var getPre func(i, j byte) int
	getPre = func(i, j byte) int {
		if i >= j {
			return prePresum[i] - prePresum[j]
		}
		return getPre(i, 0) + getPre(26, j) + previousCost[0]
	}
	var ret int
	for i := range s {
		ret += min(getNext(s[i]-'a', t[i]-'a'), getPre(s[i]-'a', t[i]-'a'))
	}
	return int64(ret)
}
