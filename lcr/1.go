package main

import "math"

func divide(a int, b int) int {
	var flag bool
	if a < 0 {
		flag = !flag
		a = -a
	}
	if b < 0 {
		flag = !flag
		b = -b
	}
	var ret int
	for a >= b {
		cnt := 1
		tmp := b
		for a > tmp+tmp {
			tmp += tmp
			cnt += cnt
		}
		a -= tmp
		ret += cnt
	}
	if flag {
		ret = -ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return math.MaxInt32
	}
	return ret
}
