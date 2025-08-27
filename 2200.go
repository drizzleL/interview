package main

func findKDistantIndices(nums []int, key int, k int) []int {
	var ret []int
	seen := make([]bool, len(nums))
	for i, j := 0, -1; i < len(nums); i++ {
		if nums[i] == key {
			j = i + k
		}
		if i <= j {
			seen[i] = true
		}
	}
	for i, j := len(nums)-1, len(nums); i >= 0; i-- {
		if nums[i] == key {
			j = i - k
		}
		if i >= j {
			seen[i] = true
		}
	}
	for i, v := range seen {
		if v {
			ret = append(ret, i)
		}
	}
	return ret
}
