package main

func longestEqualSubarray(nums []int, k int) int {
	pos := map[int][]int{}
	for i, num := range nums {
		pos[num] = append(pos[num], i)
	}
	var ret int
	used := map[int]int{}
	for num, v := range pos {
		var i int
		for j := 0; j < len(v); j++ {
			if j != 0 {
				used[num] += v[j] - v[j-1] - 1
			}
			for used[num] > k {
				used[num] -= v[i+1] - v[i] - 1
				i++
			}
			ret = max(ret, j-i+1)
		}
	}
	return ret
}
