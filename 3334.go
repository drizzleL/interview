package main

func maxScore3(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0] * nums[0])
	}
	suffixGcds := make([]int, len(nums))
	suffixGcds[len(nums)-1] = nums[len(nums)-1]
	suffixLcms := make([]int, len(nums))
	suffixLcms[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffixGcds[i] = gcd(suffixGcds[i+1], nums[i])
		suffixLcms[i] = lcm(suffixLcms[i+1], nums[i])
	}
	ret := max(suffixGcds[0]*suffixLcms[0], suffixGcds[1]*suffixLcms[1])
	gcdVal, lcmVal := nums[0], nums[0]
	for i := 1; i < len(nums)-1; i++ {
		ret = max(ret, gcd(gcdVal, suffixGcds[i+1])*lcm(lcmVal, suffixLcms[i+1]))
		gcdVal = gcd(gcdVal, nums[i])
		lcmVal = lcm(lcmVal, nums[i])
	}
	ret = max(ret, gcdVal*lcmVal)
	return int64(ret)
}

func lcm(a, b int) int {
	v := gcd(a, b)
	return a / v * b
}
