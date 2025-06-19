package main

import (
	"fmt"
)

func minOperations9(n int) int {
	str := fmt.Sprintf("%b", n)
	var ret int
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			continue
		}
		if i+1 == len(str) || str[i+1] == '0' {
			ret += 1
			continue
		}
		var cnt int
		for {
			for i < len(str) && str[i] == '1' {
				cnt += 1
				i += 1
			}
			if i == len(str) || i+1 == len(str) {
				break
			}
			if str[i+1] == '0' {
				break
			} else {
				ret += 1
			}
			i++
		}
		ret += 2
	}
	return ret
}
