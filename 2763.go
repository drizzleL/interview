package main

func sumImbalanceNumbers(nums []int) int {
	var ret int
	for i := 0; i < len(nums)-1; i++ {
		minVal := nums[i]
		maxVal := nums[i]
		dict := map[int]bool{}
		dict[nums[i]] = true
		var curr int
		for j := i + 1; j < len(nums); j++ {
			if dict[nums[j]] { // used before
				ret += curr
				continue
			}
			switch {
			case nums[j] > maxVal:
				if !dict[nums[j]-1] {
					curr += 1
				}
			case nums[j] < minVal:
				if !dict[nums[j]+1] {
					curr += 1
				}
			case dict[nums[j]-1] && dict[nums[j]+1]:
				curr -= 1
			default:
				if !dict[nums[j]+1] && !dict[nums[j]-1] {
					curr += 1
				}
			}
			minVal = min(minVal, nums[j])
			maxVal = max(maxVal, nums[j])
			dict[nums[j]] = true
			ret += curr
		}
	}
	return ret
}
