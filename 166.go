package main

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var flag int
	if numerator < 0 {
		flag ^= 1
		numerator = -numerator
	}
	if denominator < 0 {
		flag ^= 1
		denominator = -denominator
	}
	a := numerator / denominator
	numerator %= denominator
	var nums []int
	dict := map[int]int{}
	joinInts := func(nums []int) string {
		var strs []string
		for _, num := range nums {
			strs = append(strs, strconv.Itoa(num))
		}
		return strings.Join(strs, "")
	}
	ret := strconv.Itoa(a)
	if flag == 1 {
		ret = "-" + ret
	}
	for numerator != 0 {
		if idx, ok := dict[numerator]; ok { // seen before, break
			ret += "." + joinInts(nums[:idx]) + "(" + joinInts(nums[idx:]) + ")"
			return ret
		}
		dict[numerator] = len(nums)
		numerator *= 10
		nums = append(nums, numerator/denominator)
		numerator %= denominator
	}
	if len(nums) > 1 {
		ret += "." + joinInts(nums)
	}
	return ret
}
