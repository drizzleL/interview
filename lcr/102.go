package main

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	dict := [2001]int{}
	dict[1000] = 1
	for i := 0; i < len(nums); i++ {
		newdict := [2001]int{}
		for k, v := range dict {
			if v == 0 {
				continue
			}
			newdict[k+nums[i]] += v
			newdict[k-nums[i]] += v
		}
		dict = newdict
	}
	return dict[target+1000]
}
