package main

import "log"

func countCompleteSubarrays(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	cnt := len(dict)
	dict = map[int]int{}
	var ret int
	var curr int
	for i, j := 0, 0; i < len(nums); i++ {
		for ; j < len(nums) && curr != cnt; j++ {
			if dict[nums[j]] == 0 {
				curr += 1
			}
			dict[nums[j]] += 1
		}
		if curr != cnt {
			break
		}
		log.Println(j)
		ret += len(nums) - j + 1
		if dict[nums[i]] == 1 {
			curr -= 1
		}
		dict[nums[i]] -= 1
	}
	return ret
}
