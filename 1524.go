package main

func numOfSubarrays2(arr []int) int {
	var ret int
	var preOdd, preEven int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			ret += preOdd
			preEven += 1
		} else {
			ret += 1
			ret += preEven
			preOdd, preEven = preEven+1, preOdd
		}
		preOdd %= 1e9 + 7
		preEven %= 1e9 + 7
		ret %= 1e9 + 7
	}
	return ret
}
