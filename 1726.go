package main

func tupleSameProduct(nums []int) int {
	dict := map[int]int{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			dict[nums[i]*nums[j]] += 1
		}
	}
	var ret int
	for _, v := range dict {
		ret += combination(v, 2)
	}
	return ret
}
