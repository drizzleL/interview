package main

import "sort"

func maxHeight(cuboids [][]int) int {
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][2] != cuboids[j][2] {
			return cuboids[i][2] > cuboids[j][2]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] > cuboids[j][1]
		}
		return cuboids[i][0] > cuboids[j][0]
	})
	cache := map[[2]int]int{}
	var helper func(i int, pre int) int
	helper = func(i int, pre int) (ret int) {
		if i == len(cuboids) {
			return 0
		}
		if pre == -1 {
			return max(helper(i+1, pre), cuboids[i][2]+helper(i+1, i))
		}
		if c, ok := cache[[2]int{i, pre}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, pre}] = ret
		}()
		ret = helper(i+1, pre)
		if cuboids[i][1] <= cuboids[pre][1] && cuboids[i][0] <= cuboids[pre][0] {
			ret = max(ret, helper(i+1, i)+cuboids[i][2])
		}
		return
	}
	return helper(0, -1)
}
