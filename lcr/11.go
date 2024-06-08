package main

func findMaxLength(nums []int) int {
	var presum int
	var ret int
	dict := map[int]int{ // key:presum, val:firstidx
		0: -1,
	}
	for i, num := range nums {
		presum += num*2 - 1 // 0=>-1,1=>1
		if j, ok := dict[presum]; !ok {
			dict[presum] = i
		} else {
			if i-j > ret {
				ret = i - j
			}
		}
	}
	return ret
}
