package main

func countStableSubarrays2(nums []int, queries [][]int) []int64 {
	var gaps [][2]int
	curr := [2]int{0, 0}
	dict := make([]int, len(nums))
	dict[0] = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			dict[i] = len(gaps)
			curr[1] = i
			continue
		}
		gaps = append(gaps, curr)
		dict[i] = len(gaps)
		curr = [2]int{i, i}
	}
	gaps = append(gaps, curr)

	helper := func(i, j int) int64 {
		size := j - i + 1
		return int64(size * (size + 1) / 2)
	}
	var gapsSum []int64
	var sum int64
	for _, g := range gaps {
		sum += helper(g[0], g[1])
		gapsSum = append(gapsSum, sum)
	}
	gSum := func(idx1, idx2 int) int64 {
		if idx1 > idx2 {
			return 0
		}
		if idx1 == 0 {
			return gapsSum[idx2]
		}
		return gapsSum[idx2] - gapsSum[idx1-1]
	}
	ret := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		idx1, idx2 := dict[l], dict[r]
		if idx1 == idx2 { // belong same gap
			ret[i] = helper(l, r)
			continue
		}
		g1, g2 := gaps[idx1], gaps[idx2]
		ret[i] = gSum(idx1+1, idx2-1) + helper(l, g1[1]) + helper(g2[0], r)
	}
	return ret
}
