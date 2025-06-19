package main

func isArraySpecial(nums []int, queries [][]int) []bool {
	arr := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = arr[i-1]
		}
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		ret[i] = arr[q[0]] == arr[q[1]]
	}
	return ret
}
