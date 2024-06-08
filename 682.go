package main

import "strconv"

func calPoints(operations []string) int {
	var scores []int
	for _, op := range operations {
		switch op {
		case "+":
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		case "D":
			scores = append(scores, scores[len(scores)-1]*2)
		case "C":
			scores = scores[:len(scores)-1]
		default:
			num, _ := strconv.Atoi(op)
			scores = append(scores, num)
		}
	}
	var ret int
	for _, c := range scores {
		ret += c
	}
	return ret
}
