package main

func hasSameDigits2(s string) bool {
	cal := func(a int, mod int) (int, int) {
		count := 0
		for a > 0 && a%mod == 0 {
			count += 1
			a /= mod
		}
		return a % mod, count
	}
	test := func(mod int) bool {
		n := len(s)
		res := 0
		r := 1
		c := 0
		for i := 0; i < n-1; i++ {
			if c == 0 {
				res += r * (int(s[i]) - int(s[i+1]))
			}
			rr, cc := cal(n-2-i, mod)
			r = r * rr % mod
			c += cc

			rr, cc = cal(i+1, mod)
			r = r * pow(rr, mod-2, mod) % mod
			c -= cc
		}
		return res%mod == 0
	}
	return test(2) && test(5)
}
