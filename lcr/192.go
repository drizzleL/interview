package main

import "math"

func myAtoi(str string) int {
	var i int
	for i = 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		break
	}
	if i == len(str) {
		return 0
	}
	var flag bool
	switch str[i] {
	case '-':
		flag = true
		fallthrough
	case '+':
		i++
	}
	var v int
	for ; i < len(str); i++ {
		if str[i] > '9' || str[i] < '0' {
			break
		}
		v = v*10 + int(str[i]-'0')
	}
	if flag {
		v = -v
	}
	if v > math.MaxInt32 {
		return math.MaxInt32
	}
	if v < math.MinInt32 {
		return math.MinInt32
	}
	return v
}
