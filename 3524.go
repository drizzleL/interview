package main

func resultArray(nums []int, k int) []int64 {
	ret := make([]int64, k)
	dict := make([]int, k) // remainder accum
	for i := 0; i < len(nums); i++ {
		newdict := make([]int, k)
		for j := 0; j < k; j++ {
			newr := (nums[i] * j) % k
			ret[newr] += int64(dict[j])
			newdict[newr] += dict[j]
		}
		dict = newdict
		ret[nums[i]%k] += 1
		dict[nums[i]%k] += 1
	}
	return ret
}
