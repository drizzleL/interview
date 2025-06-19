package main

func isPossible(nums []int) bool {
	head, tail := map[int]int{}, map[int]int{}
	for _, num := range nums {
		head[num] += 1
	}
	for _, num := range nums {
		if head[num] == 0 {
			continue
		}
		head[num] -= 1
		if tail[num] > 0 {
			tail[num] -= 1
			tail[num+1] += 1
			continue
		}
		if head[num+1] > 0 && head[num+2] > 0 {
			head[num+1] -= 1
			head[num+2] -= 1
			tail[num+3] += 1
			continue
		}
		return false
	}
	return true
}
