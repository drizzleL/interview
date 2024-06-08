package main

func longestConsecutive(nums []int) int {
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
	}
	var ret int
	for _, num := range nums {
		if dict[num-1] {
			continue
		}
		size := 1
		for dict[num+1] {
			size++
			num = num + 1
		}
		if size > ret {
			ret = size
		}
	}
	return ret
}
