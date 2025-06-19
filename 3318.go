package main

import "sort"

func findXSum(nums []int, k int, x int) []int {
	dict := map[int]int{}
	for i := 0; i < k; i++ {
		dict[nums[i]] += 1
	}
	check := func() int {
		var tmp [][2]int
		for k, v := range dict {
			tmp = append(tmp, [2]int{k, v})
		}
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i][1] == tmp[j][1] {
				return tmp[i][0] > tmp[j][0]
			}
			return tmp[i][1] > tmp[j][1]
		})
		var ret int
		for i := 0; i < x && i < len(tmp); i++ {
			ret += tmp[i][0] * tmp[i][1]
		}
		return ret
	}
	ret := make([]int, 0, len(nums)-k+1)
	ret = append(ret, check())
	for i, j := k, 0; i < len(nums); i, j = i+1, j+1 {
		dict[nums[i-k]] -= 1
		dict[nums[i]] += 1
		ret = append(ret, check())
	}
	return ret
}
