package main

func countValidSelections(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	var leftSum int
	var ret int
	for _, num := range nums {
		leftSum += num
		if num != 0 {
			continue
		}
		switch leftSum * 2 {
		case sum:
			ret += 2
		case sum - 1, sum + 1:
			ret += 1
		}
	}
	return ret
}
