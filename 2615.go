package main

func distance(nums []int) []int64 {
	leftDict, rightDict := map[int]int{}, map[int]int{}
	leftCnt, rightCnt := map[int]int{}, map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		leftDict[num] += i
		leftCnt[num] += 1
	}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		rightDict[num] += len(nums) - 1 - i
		rightCnt[num] += 1
	}
	ret := make([]int64, len(nums))
	last := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		leftDict[num] -= (i - last[num]) * leftCnt[num]
		ret[i] += int64(leftDict[num])
		leftCnt[num] -= 1
		last[num] = i
	}
	last = map[int]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if _, ok := last[num]; !ok {
			last[num] = len(nums) - 1
		}
		rightDict[num] -= (last[num] - i) * rightCnt[num]
		ret[i] += int64(rightDict[num])
		rightCnt[num] -= 1
		last[num] = i
	}
	return ret
}
