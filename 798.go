package main

func bestRotation(nums []int) int {
	change := make([]int, len(nums))
	for i, v := range nums {
		change[(i-v+1+len(nums))%len(nums)] -= 1
	}
	var maxI int
	for i := 1; i < len(nums); i++ {
		change[i+1] += change[i] + 1
		if change[i] > change[maxI] {
			maxI = i
		}
	}
	return maxI
}
