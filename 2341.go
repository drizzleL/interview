package main

func numberOfPairs(nums []int) []int {
	ret := []int{0, 0}
	dict := map[int]bool{}
	for _, num := range nums {
		if dict[num] { // found pair
			delete(dict, num)
			ret[0] += 1
			continue
		}
		dict[num] = true
	}
	ret[1] = len(dict)
	return ret
}
