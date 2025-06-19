package main

import "log"

func numberOfSubsequences(nums []int) int64 {
	dict2 := map[[2]int]int{}
	getK := func(a, b int) [2]int {
		g := gcd(a, b)
		return [2]int{a / g, b / g}
	}
	for i := len(nums) - 1; i >= 6; i-- {
		for j := i - 2; j >= 4; j-- {
			dict2[getK(nums[i], nums[j])] += 1
		}
	}
	var ret int
	for q := 2; q <= len(nums)-5; q++ {
		log.Println(dict2, ret)
		for p := 0; p < q-1; p++ {
			key := getK(nums[p], nums[q])
			ret += dict2[key]
		}
		for s := len(nums) - 1; s >= q+4; s-- {
			key := getK(nums[s], nums[q+2])
			dict2[key] -= 1
		}
	}
	return int64(ret)
}
