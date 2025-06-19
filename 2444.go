package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	if len(nums) == 0 {
		return 0
	}
	currMin, currMax := nums[0], nums[0]
	var rightMost int
	for ; rightMost < len(nums); rightMost++ {
		if nums[rightMost] > maxK || nums[rightMost] < minK {
			break
		}
		currMax = max(currMax, nums[rightMost])
		currMin = min(currMin, nums[rightMost])
	}
	var ret int64
	if rightMost != len(nums) { // found weird num
		ret = countSubarrays(nums[rightMost+1:], minK, maxK)
	}
	if currMin != minK || currMax != maxK {
		return ret
	}
	leftMost := 0
	lastMin, lastMax := -1, -1
	for i := 0; i < rightMost; i++ {
		if nums[i] == minK {
			lastMin = i
		}
		if nums[i] == maxK {
			lastMax = i
		}
		if lastMin == -1 || lastMax == -1 {
			continue
		}
		ret += int64((min(lastMin, lastMax) - leftMost + 1) * (rightMost - max(lastMin, lastMax)))
		leftMost = min(lastMin, lastMax) + 1
		lastMin, lastMax = -1, -1 // reset again
		if nums[i] == minK {
			lastMin = i
		}
		if nums[i] == maxK {
			lastMax = i
		}
	}
	return ret
}
