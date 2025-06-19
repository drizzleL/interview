package main

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i] = 1
		if i != 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		right[i] = 1
		if i != len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	var ret int
	for i := 0; i < len(ratings); i++ {
		ret += max(left[i], right[i])
	}
	return ret
}
