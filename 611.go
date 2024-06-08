package main

import (
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var ret int
	for i := 0; i < len(nums)-2; i++ {
		// if i != 0 && nums[i-1] == nums[i] {
		// 	continue
		// }
		for j := i + 1; j < len(nums)-1; j++ {
			// if j != i+1 && nums[j-1] == nums[j] {
			// 	continue
			// }
			s := nums[i] + nums[j]
			idx := sort.Search(len(nums)-j-1, func(i int) bool {
				return nums[j+1+i] >= s
			})
			ret += idx
		}
	}
	return ret
}
