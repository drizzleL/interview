package main

func maxChunksToSorted(arr []int) int {
	var ret int
	var start int
	flag := make([]bool, len(arr))
	for i, num := range arr {
		flag[num] = true
		var up bool
		for j := start; j <= i; j++ {
			if !flag[j] {
				up = true
			}
		}
		if !up {
			start = i + 1
			ret += 1
		}
	}
	return ret
}
