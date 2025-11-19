package main

func splitArray77(nums []int) int64 {
	if len(nums) == 2 {
		return int64(abs(nums[0] - nums[1]))
	}
	left, right := 0, len(nums)-1
	var leftSum, rightSum int
	for {
		leftSum += nums[left]
		if left+1 < len(nums) && nums[left+1] > nums[left] {
			left += 1
			continue
		}
		break
	}
	if left == len(nums) { // all increasing
		return -1
	}
	for {
		rightSum += nums[right]
		if right-1 >= 0 && nums[right-1] > nums[right] {
			right -= 1
			continue
		}
		break
	}
	if left == right {
		sum1 := abs(leftSum - rightSum + nums[left])
		sum2 := abs(rightSum - leftSum + nums[left])
		return int64(min(sum1, sum2))
	}
	if left+1 == right {
		return int64(abs(leftSum - rightSum))
	}
	return -1
}
