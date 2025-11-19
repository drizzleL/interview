package main

func countMajoritySubarrays2(nums []int, target int) int64 {
	var cnt, ret int
	f := NewFenwick(len(nums)*2 + 1)
	f.add(len(nums)+1, 1)
	for i, num := range nums {
		if num == target {
			cnt += 1
		}
		key := cnt*2 - i + len(nums)
		ret += f.sum(key - 1)
		f.add(key, 1)
	}
	return int64(ret)
}
