package main

func numTriplets(nums1 []int, nums2 []int) int {
	helper := func(nums1, nums2 []int) int {
		var ret int
		dict := map[int]int{}
		for _, num := range nums2 {
			for _, num2 := range nums1 {
				v := num2 * num2
				if v/num*num == v {
					ret += dict[v/num]
				}
			}
			dict[num] += 1
		}
		return ret
	}
	return helper(nums1, nums2) + helper(nums2, nums1)
}
