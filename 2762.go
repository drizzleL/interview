package main

func continuousSubarrays(nums []int) int64 {
	var maxs, mins []int
	var ret int
	var last int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for (len(mins) > 0 && num > nums[mins[0]]+2) || (len(maxs) > 0 && num < nums[maxs[0]]-2) {
			if last == maxs[0] {
				maxs = maxs[1:]
			}
			if last == mins[0] {
				mins = mins[1:]
			}
			last += 1
		}
		ret += i - last + 1
		for len(maxs) != 0 && nums[i] > nums[maxs[len(maxs)-1]] {
			maxs = maxs[:len(maxs)-1]
		}
		maxs = append(maxs, i)
		for len(mins) != 0 && nums[i] < nums[mins[len(mins)-1]] {
			mins = mins[:len(mins)-1]
		}
		mins = append(mins, i)
	}
	return int64(ret)
}
