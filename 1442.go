package main

func countTriplets(arr []int) int {
	var ret int
	for j := 1; j < len(arr); j++ {
		dict := map[int]int{}
		var prefixSum int
		for i := j - 1; i >= 0; i-- {
			prefixSum ^= arr[i]
			dict[prefixSum] += 1
		}
		var suffixSum int
		for k := j; k < len(arr); k++ {
			suffixSum ^= arr[k]
			ret += dict[suffixSum]
		}
	}
	return ret
}
