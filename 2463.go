package main

import (
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	cache := map[[2]int]int{}
	var helper func(robotIdx, factoryIdx int) int
	helper = func(robotIdx, factoryIdx int) (ret int) {
		if robotIdx == len(robot) {
			return 0
		}
		if factoryIdx == len(factory) {
			return -1
		}
		if c, ok := cache[[2]int{robotIdx, factoryIdx}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{robotIdx, factoryIdx}] = ret
		}()
		ret = helper(robotIdx, factoryIdx+1)
		f := factory[factoryIdx]
		var dis int
		for i := 0; i < f[1] && robotIdx+i < len(robot); i++ {
			dis += abs(f[0] - robot[robotIdx+i])
			if v := helper(robotIdx+i+1, factoryIdx+1); v != -1 {
				if ret == -1 || dis+v < ret {
					ret = dis + v
				}
			}
		}
		return
	}
	return int64(helper(0, 0))
}
