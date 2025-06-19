package main

func maximumSubarraySum(nums []int, k int) int64 {
	var ret int
	var sum int
	dict := map[int]int{}
	var multi int
	for i := 0; i < k; i++ {
		sum += nums[i]
		dict[nums[i]] += 1
		if dict[nums[i]] == 2 {
			multi += 1
		}
	}
	if multi == 0 {
		ret = sum
	}
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		dict[nums[i-k]] -= 1
		if dict[nums[i-k]] == 1 {
			multi -= 1
		}
		sum += nums[i]
		dict[nums[i]] += 1
		if dict[nums[i]] == 2 {
			multi += 1
		}
		if multi == 0 {
			ret = max(ret, sum)
		}
	}
	return int64(ret)
}
