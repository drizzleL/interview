package main

import "strconv"

func minMaxDifference(num int) int {
	str := []byte(strconv.Itoa(num))
	helper := func(from, to byte) int {
		var b []byte
		for i := 0; i < len(str); i++ {
			if str[i] == from {
				b = append(b, to)
			} else {
				b = append(b, str[i])
			}
		}
		v, _ := strconv.Atoi(string(b))
		return v
	}
	var firstMin int
	for i := 0; i < len(str); i++ {
		if str[i] != '9' {
			firstMin = i
			break
		}
	}
	return helper(str[firstMin], '9') - helper(str[0], '0')
}
