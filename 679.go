package main

func judgePoint24(cards []int) bool {
	eps := 1e-6
	checkZero := func(v float64) bool {
		return v < eps && v > -eps
	}
	compute := func(a, b float64) (ret []float64) {
		ret = append(ret, a+b, a-b, b-a, a*b)
		if !checkZero(b) {
			ret = append(ret, a/b)
		}
		if !checkZero(a) {
			ret = append(ret, b/a)
		}
		return
	}
	var dfs func(nums []float64) bool
	dfs = func(nums []float64) bool {
		if len(nums) == 1 {
			return checkZero(nums[0] - 24)
		}
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				var next []float64
				for k := 0; k < len(nums); k++ {
					if k == i || k == j {
						continue
					}
					next = append(next, nums[k])
				}
				for _, num := range compute(nums[i], nums[j]) {
					next = append(next, num)
					if dfs(next) {
						return true
					}
					next = next[:len(next)-1]
				}
			}
		}
		return false
	}
	var nums []float64
	for _, card := range cards {
		nums = append(nums, float64(card))
	}
	return dfs(nums)
}
