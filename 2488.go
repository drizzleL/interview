package main

func countSubarrays6(nums []int, k int) int {
	dict := map[int]int{}
	var i int    // where nums[i] == k
	var curr int // current gap
	for ; nums[i] != k; i++ {
	}
	for j := i; j >= 0; j-- {
		switch {
		case nums[j] > k:
			curr += 1
		case nums[j] < k:
			curr -= 1
		}
		dict[curr] += 1
	}
	var ret int
	curr = 0 // reset curr
	for j := i; j < len(nums); j++ {
		switch {
		case nums[j] > k:
			curr += 1
		case nums[j] < k:
			curr -= 1
		}
		ret += dict[-curr]
		ret += dict[-curr+1]
	}
	return ret
}
