package main

func nthUglyNumber(n int) int {
	var a, b, c int
	nums := []int{1}
	for i := 1; i < n; i++ {
		va, vb, vc := nums[a]*2, nums[b]*3, nums[c]*5

		v := min(va, min(vb, vc))
		nums = append(nums, v)
		if v == va {
			a++
		}
		if v == vb {
			b++
		}
		if v == vc {
			c++
		}
	}
	return nums[n-1]
}
