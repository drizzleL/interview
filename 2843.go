package main

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	var ret int
	for i := low; i <= high; i++ {
		str := strconv.Itoa(i)
		if len(str)%2 == 1 {
			continue
		}
		var cnt int
		for j, k := 0, len(str)-1; j < k; j, k = j+1, k-1 {
			cnt += int(str[j] - '0')
			cnt -= int(str[k] - '0')
		}
		if cnt == 0 {
			ret += 1
		}
	}
	return ret
}
