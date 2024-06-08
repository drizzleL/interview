package main

func findPrefixScore(nums []int) []int64 {
	var ret []int64
	var presum int
	var premax int
	for _, num := range nums {
		premax = max(premax, num)
		presum += premax + num
		ret = append(ret, int64(presum))
	}
	return ret
}
