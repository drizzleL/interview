package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func combinationSum42(nums []int, target int) int {
	var helper func(val int) int
	cache := map[int]int{}
	helper = func(val int) (ret int) {
		if c, ok := cache[val]; ok {
			return c
		}
		defer func() {
			cache[val] = ret
		}()
		if val == target {
			return 1
		}
		for _, num := range nums {
			if val+num > target {
				continue
			}
			ret += helper(val + num)
		}
		return ret
	}
	return helper(0)
}
