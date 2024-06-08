package main

func sumFourDivisors(nums []int) int {
	var maxV int
	for _, num := range nums {
		maxV = max(maxV, num)
	}
	dict := make([]int, maxV+1)
	times := make([]int, maxV+1)
	for i := 1; i <= maxV; i++ {
		for j := 1; i*j <= maxV; j++ {
			dict[i*j] += i
			times[i*j] += 1
		}
	}
	var ret int
	for _, num := range nums {
		if times[num] == 4 {
			ret += dict[num]
		}
	}
	return ret
}
