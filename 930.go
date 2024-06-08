package main

func numSubarraysWithSum(nums []int, goal int) int {
	var presum int
	var ret int
	dict := map[int]int{
		0: 1,
	}
	for _, num := range nums {
		presum += num
		ret += dict[presum-goal]
		dict[presum] += 1
	}
	return ret
}
