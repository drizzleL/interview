package main

func countGood(nums []int, k int) int64 {
	var ret, pairs int
	dict := map[int]int{}
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && pairs < k {
			pairs -= combination(dict[nums[j]], 2)
			dict[nums[j]] += 1
			pairs += combination(dict[nums[j]], 2)
			j += 1
		}
		if pairs < k {
			break
		}
		ret += len(nums) - j + 1
		pairs -= combination(dict[nums[i]], 2)
		dict[nums[i]] -= 1
		pairs += combination(dict[nums[i]], 2)
	}
	return int64(ret)
}
