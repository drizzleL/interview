package main

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	helper := func(x int) bool {
		s := strconv.Itoa(x)
		return !strings.Contains(s, "0")
	}
	for i := 1; i < n; i++ {
		if helper(i) && helper(n-i) {
			return []int{i, n - i}
		}
	}
	return nil
}
