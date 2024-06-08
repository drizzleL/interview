package main

func numWaterBottles(numBottles int, numExchange int) int {
	var ret int
	var empty int
	for numBottles > 0 {
		ret += numBottles
		empty += numBottles
		numBottles = empty / numExchange
		empty -= numBottles
	}
	return ret
}
