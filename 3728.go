package main

func countStableSubarrays(capacity []int) int64 {
	var ret int
	dict := map[[2]int]int{}
	var pre1, pre2 int
	pre2 = capacity[0]
	for i := 1; i < len(capacity); i++ {
		v := capacity[i]
		pre2 += v
		key := [2]int{pre2 - v*2, v}
		ret += dict[key]

		pre1 += capacity[i-1]
		dict[[2]int{pre1, capacity[i-1]}] += 1
	}
	return int64(ret)
}
