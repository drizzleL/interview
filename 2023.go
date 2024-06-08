package main

func numOfPairs(nums []string, target string) int {
	targetSize := len(target)
	dict := map[string]int{}
	var ret int
	for _, num := range nums {
		numSize := len(num)
		if numSize >= targetSize {
			continue
		}
		if target[:numSize] == num {
			ret += dict[target[numSize:]]
		}
		if target[targetSize-numSize:] == num {
			ret += dict[target[:targetSize-numSize]]
		}
		dict[num] += 1
	}
	return ret
}
