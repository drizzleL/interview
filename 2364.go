package main

func countBadPairs(nums []int) int64 {
	var goodPair int
	dict := map[int]int{}
	for i, num := range nums {
		goodPair += dict[num-i]
		dict[num-i] += 1
	}
	return int64((len(nums)-1)*len(nums)/2 - goodPair)
}
