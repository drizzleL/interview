package main

import "sort"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var diffs []int
	var ret int
	for _, num := range nums {
		ret += num
		diffs = append(diffs, num^k-num)
	}
	sort.Ints(diffs)
	for i := len(diffs) - 2; i >= 0; i -= 2 {
		tmp := diffs[i] + diffs[i+1]
		if tmp > 0 {
			ret += tmp
			continue
		}
		break
	}
	return int64(ret)
}

func maximumValueSum2(nums []int, k int, edges [][]int) int64 {
	var ret int
	for _, num := range nums {
		ret += num
	}
	n := len(nums) + 1
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	dirDict := make([][]int, n)
	seen := make([]bool, n)
	var path func(i int)
	path = func(i int) {
		seen[i] = true
		for _, next := range dict[i] {
			if seen[next] {
				continue
			}
			dirDict[i] = append(dirDict[i], next)
			path(next)
		}
	}
	path(0)
	cache := map[[2]int]int{}
	flip := func(a, b int, flag int) int {
		if flag == 1 {
			a ^= k
		}
		var ret int
		ret += a ^ k - a
		ret += b ^ k - b
		return ret
	}
	var helper func(i int, flag int) int
	helper = func(i int, flag int) (ret int) {
		if c, ok := cache[[2]int{i, flag}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, flag}] = ret
		}()
		dirs := dirDict[i]
		if len(dirs) == 0 {
			return 0
		}
		if len(dirs) == 1 {
			return max(helper(dirs[0], 0), flip(nums[i], nums[dirs[0]], flag)+helper(dirs[0], 1))
		}
		ret = helper(dirs[0], 0) + helper(dirs[1], 0)
		ret = max(ret, flip(nums[i], nums[dirs[0]], flag)+helper(dirs[0], 1)+flip(nums[i], nums[dirs[1]], flag)+helper(dirs[1], 1))
		ret = max(ret, helper(dirs[0], 0)+flip(nums[i], nums[dirs[1]], flag)+helper(dirs[1], 1))
		ret = max(ret, helper(dirs[1], 0)+flip(nums[i], nums[dirs[0]], flag)+helper(dirs[0], 1))
		return
	}
	ret += helper(0, 0)
	return int64(ret)
}
