package main

func maxSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	dict := map[int]int{}
	var ret int
	var last int
	for i, num := range nums {
		dict[num] += 1
		if dict[num] <= k {
			ret = max(ret, i-last+1)
			continue
		}
		for {
			v := nums[last]
			dict[v] -= 1
			last += 1
			if v == num {
				break
			}
		}
	}
	return ret
}
