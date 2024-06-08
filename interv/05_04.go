package main

func findClosedNumbers(num int) []int {
	isZero := func(val int, pos int) bool {
		return (1<<pos)&val == 0
	}
	swap := func(num int, posZero, posOne int) int {
		num |= 1 << posOne
		num &= ^(1 << posZero)
		return num
	}
	var ret []int
	// find 01
	{
		var one int
		for ; one <= 29; one++ {
			if !isZero(num, one) && isZero(num, one+1) {
				break
			}
		}
		if one <= 29 {
			val := swap(num, one, one+1)
			for i, j := one-1, 0; i > j && !isZero(val, i) && isZero(val, j); i, j = i-1, j+1 {
				val = swap(val, i, j)
			}
			ret = append(ret, val)
		} else {
			ret = append(ret, -1)
		}
	}
	// find 10
	{
		var zero int
		for ; zero <= 29; zero++ {
			if num&(1<<zero) == 0 && num&(1<<(zero+1)) != 0 {
				break
			}
		}
		if zero <= 29 {
			val := swap(num, zero+1, zero)
			for i, j := zero-1, 0; i > j && isZero(val, i) && !isZero(val, j); i, j = i-1, j+1 {
				val = swap(val, j, i)
			}
			ret = append(ret, val)
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}
