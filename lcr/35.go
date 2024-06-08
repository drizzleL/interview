package main

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) <= 1 {
		return 0
	}
	toDigit := func(s string) int {
		ss := strings.Split(s, ":")
		h, _ := strconv.Atoi(ss[0])
		m, _ := strconv.Atoi(ss[1])
		return h*60 + m
	}
	sort.Slice(timePoints, func(i, j int) bool {
		a, b := timePoints[i], timePoints[j]
		return toDigit(a) < toDigit(b)
	})
	ret := toDigit(timePoints[0]) - toDigit(timePoints[len(timePoints)-1]) + 24*60
	for i := 0; i < len(timePoints)-1; i++ {
		tmp := toDigit(timePoints[i+1]) - toDigit(timePoints[i])
		if tmp < ret {
			ret = tmp
		}
	}
	return ret
}
