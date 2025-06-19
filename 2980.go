package main

func hasTrailingZeros(nums []int) bool {
	var flag bool
	for _, num := range nums {
		if num%2 == 1 {
			continue
		}
		if flag {
			return true
		}
		flag = true
	}
	return false
}
