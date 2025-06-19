package main

func countLargestGroup(n int) int {
	var maxSize int
	dict := map[int]int{}
	for i := 1; i <= n; i++ {
		k := i
		var idx int
		for k != 0 {
			idx += k % 10
			k /= 10
		}
		dict[idx] += 1
		maxSize = max(maxSize, dict[idx])
	}
	var ret int
	for _, v := range dict {
		if v == maxSize {
			ret += 1
		}
	}
	return ret
}
