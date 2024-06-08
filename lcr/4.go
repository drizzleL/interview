package main

func singleNumber(nums []int) int {
	var data [32]int
	var flag int
	for _, num := range nums {
		if num < 0 {
			num = -num
			flag += 1
		}
		tmp := 1
		for i := range data {
			if tmp&num != 0 {
				data[i] += 1
			}
			if data[i] == 3 {
				data[i] = 0
			}
			tmp *= 2
		}
	}
	var ret int
	tmp := 1
	for i := range data {
		if data[i] != 0 {
			ret += tmp
		}
		tmp *= 2
	}
	if flag%3 != 0 {
		ret = -ret
	}
	return ret
}
