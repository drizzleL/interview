package main

func subarrayLCM(nums []int, k int) int {
	split := func(x int) map[int]int {
		dict := map[int]int{}
		for i := 2; x != 1; i++ {
			for x%i == 0 {
				dict[i] += 1
				x /= i
			}
		}
		return dict
	}
	baseDict := split(k)
	check := func(numDict map[int]int) bool {
		for k, v := range numDict {
			if v > baseDict[k] {
				return false
			}
		}
		return true
	}
	var ret int
	for i := range nums {
		dict := map[int]bool{}
		for j := i; j < len(nums); j++ {
			numDict := split(nums[j])
			if !check(numDict) {
				break
			}
			for k, v := range numDict {
				if v == baseDict[k] {
					dict[k] = true
				}
			}
			if len(dict) == len(baseDict) {
				ret += 1
			}
		}
	}
	return ret
}
