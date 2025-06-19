package main

func minSubarray(nums []int, p int) int {
	var need int
	for _, num := range nums {
		need += num
		need %= p
	}
	dict := map[int]int{
		0: -1,
	}
	var curr int
	ret := len(nums)
	for i, num := range nums {
		curr += num
		curr %= p
		dict[curr] = i
		want := (curr - need + p) % p
		if _, ok := dict[want]; !ok {
			continue
		}
		ret = min(ret, i-dict[want])
	}
	if ret == len(nums) {
		return -1
	}
	return ret
}
