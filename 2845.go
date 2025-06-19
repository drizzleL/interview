package main

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	var cnt int
	var ret int
	dict := map[int]int{}
	dict[0] = 1
	for _, num := range nums {
		if num%modulo == k {
			cnt += 1
		}
		cnt %= modulo
		ret += dict[(cnt-k+modulo)%modulo]
		dict[cnt%modulo] += 1
	}
	return int64(ret)
}
