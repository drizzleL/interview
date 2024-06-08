package main

import (
	"strings"
)

func validNumber(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "E", "e")
	nums := strings.Split(s, "e")
	pureNum := func(str string, mark bool) bool {
		if len(str) == 1 && (str[0] == '+' || str[0] == '-') {
			return false
		}
		if mark && len(str) >= 1 && (str[0] == '+' || str[0] == '-') {
			str = str[1:]
		}
		for _, c := range str {
			if c > '9' || c < '0' {
				return false
			}
		}
		return true
	}
	if len(nums) > 2 {
		return false
	}
	if len(nums) == 2 {
		if len(nums[1]) == 0 {
			return false
		}
		if !pureNum(nums[1], true) {
			return false
		}
		if nums[1] == "-" || nums[1] == "+" {
			return false
		}
	}
	if len(nums[0]) == 0 {
		return false
	}
	if nums[0] == "." {
		return false
	}
	preNums := strings.Split(nums[0], ".")
	if len(preNums) > 2 {
		return false
	}
	if len(preNums) == 2 {
		if !pureNum(preNums[1], false) {
			return false
		}
	}
	return pureNum(preNums[0], true)
}
