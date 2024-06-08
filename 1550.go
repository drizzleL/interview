package main

func threeConsecutiveOdds(arr []int) bool {
	var odds int
	for _, num := range arr {
		if num%2 == 0 {
			odds = 0
			continue
		}
		odds += 1
		if odds >= 3 {
			return true
		}
	}
	return false
}
