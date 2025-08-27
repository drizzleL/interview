package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var minIdx, maxIdx int
	for i := indexDifference; i < len(nums); i++ {
		last := i - indexDifference
		if nums[last] > nums[maxIdx] {
			maxIdx = last
		}
		if nums[last] < nums[minIdx] {
			minIdx = last
		}
		if nums[i]-nums[minIdx] >= valueDifference {
			return []int{minIdx, i}
		}
		if nums[maxIdx]-nums[i] >= valueDifference {
			return []int{maxIdx, i}
		}
	}
	return []int{-1, -1}
}
