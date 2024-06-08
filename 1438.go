package main

func longestSubarray(nums []int, limit int) int {
	var incrs, decrs []int
	var left int
	var ret int
	for i, num := range nums {
		for (len(incrs) > 0 && num-nums[incrs[0]] > limit) || (len(decrs) > 0 && nums[decrs[0]]-num > limit) {
			if len(incrs) != 0 && nums[incrs[0]] == nums[left] {
				incrs = incrs[1:]
			}
			if len(decrs) != 0 && nums[decrs[0]] == nums[left] {
				decrs = decrs[1:]
			}
			left++
		}
		for len(incrs) != 0 && nums[incrs[len(incrs)-1]] > num {
			incrs = incrs[:len(incrs)-1]
		}
		for len(decrs) != 0 && nums[decrs[len(decrs)-1]] < num {
			decrs = decrs[:len(decrs)-1]
		}
		incrs = append(incrs, i)
		decrs = append(decrs, i)
		ret = max(ret, i-left+1)
	}
	return ret
}
