package main

func removeZeros(n int64) int64 {
	var ret int64
	for base := int64(1); n > 0; n /= 10 {
		if n%10 == 0 {
			continue
		}
		ret += (n % 10) * base
		base *= 10
	}
	return ret
}
