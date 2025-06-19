package main

func solveQueries(nums []int, queries []int) []int {
	ret := make([]int, len(queries))
	for i := range ret {
		ret[i] = -1
	}
	prev := map[int]int{}
	for i, num := range nums {
		prev[num] = i
	}
	tmp := make([]int, len(nums))
	for i := range tmp {
		tmp[i] = -1
	}
	dis := func(a, b int) int {
		if a > b {
			a, b = b, a
		}
		return min(b-a, a+len(nums)-b)
	}
	for i, num := range nums {
		if prev[num] == i {
			continue
		}
		tmp[i] = dis(i, prev[num])
		prev[num] = i
	}
	next := map[int]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		next[nums[i]] = i
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if next[nums[i]] == i {
			continue
		}
		d := dis(i, next[nums[i]])
		if tmp[i] == -1 || d < tmp[i] {
			tmp[i] = d
		}
		next[nums[i]] = i
	}
	for i, q := range queries {
		ret[i] = tmp[q]
	}
	return ret
}
