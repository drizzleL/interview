package main

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	check := func(x int) int {
		var flag bool
		for i := len(s) - 1; i >= 0; i, x = i-1, x/10 {
			d := x % 10
			c := int(s[i] - '0')
			switch {
			case d > c:
				flag = false
			case d < c:
				flag = true
			}
		}
		if flag && x > 0 { // pass but - 1
			return 1
		}
		if flag { // not pass
			return 2
		}
		return 0 // pass
	}
	var check2 func(x int) int
	check2 = func(x int) int {
		base := 1
		base2 := 1
		var ret int
		digit := []int{}
		for x2 := x; x2 != 0; x2 /= 10 {
			digit = append(digit, base)
			base *= limit + 1
			base2 *= 10
		}
		base2 /= 10
		for x2, i := x, len(digit)-1; x2 != 0; i = i - 1 {
			d := x2 / base2
			if d > limit {
				ret += digit[i] * (limit + 1)
				return ret
			}
			ret += digit[i] * d
			x2 %= base2
			base2 /= 10
		}
		ret += 1
		return ret
	}
	helper := func(x int) int {
		if x < 0 {
			return 0
		}
		ok := check(x)
		if ok == 2 {
			return 0
		}
		for i := len(s) - 1; i >= 0; i, x = i-1, x/10 {
		}
		return check2(x - ok)
	}
	return int64(helper(int(finish)) - helper(int(start)-1))
}
