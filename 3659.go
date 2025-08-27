package main

func partitionArray2(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num] += 1
	}
	gCnt := len(nums) / k
	for _, v := range cnt {
		if v > gCnt {
			return false
		}
	}
	return true
}
