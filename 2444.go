package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	minIdx, maxIdx := -1, -1
	leftIdx := 0
	var ret int
	for i, num := range nums {
		if num > maxK || num < minK {
			minIdx, maxIdx = -1, -1
			leftIdx = i + 1
			continue
		}
		if num == maxK {
			maxIdx = i
		}
		if num == minK {
			minIdx = i
		}
		if minIdx != -1 && maxIdx != -1 {
			ret += min(minIdx, maxIdx) - leftIdx + 1
		}
	}
	return int64(ret)
}
