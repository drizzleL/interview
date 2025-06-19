package main

func minOperations7(boxes string) []int {
	var left, right int
	ret := make([]int, len(boxes))
	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			ret[0] += i
			right += 1
		}
	}
	left = int(boxes[0] - '0')
	for i := 1; i < len(boxes); i++ {
		ret[i] = ret[i-1] + left - right
		if boxes[i] == '1' {
			left += 1
			right -= 1
		}
	}
	return ret
}
