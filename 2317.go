package main

func maximumXOR(nums []int) int {
	var ret int
	for i := 0; i < 32; i++ {
		var flag bool
		for _, num := range nums {
			if num&1<<i != 0 {
				flag = true
				break
			}
		}
		if flag {
			ret |= 1 << i
		}
	}
	return ret
}
