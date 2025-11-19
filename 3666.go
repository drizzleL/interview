package main

func minOperations12(s string, k int) int {
	var zero int
	for _, c := range s {
		if c == '0' {
			zero += 1
		}
	}
	if zero == 0 {
		return 0
	}
	n := len(s)
	one := n - zero
	for i := 1; i <= n; i++ {
		p := i * k
		if (p-zero)&1 != 0 {
			continue
		}
		if i&1 != 0 { // odd
			if p >= zero && p <= (zero*i+one*(i-1)) {
				return i
			}
		} else {
			if p >= zero && p <= (zero*(i-1)+one*i) {
				return i
			}
		}
	}
	return -1
}
