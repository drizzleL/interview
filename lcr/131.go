package main

func cuttingBamboo(bamboo_len int) int {
	if bamboo_len <= 3 {
		return bamboo_len - 1
	}
	var n3 int
	ret := 1
	switch bamboo_len % 3 {
	case 2:
		n3 = bamboo_len / 3
		ret = 2
	case 1:
		n3 = bamboo_len/3 - 1
		ret = 4
	case 0:
		n3 = bamboo_len / 3
	}
	return ret * int(pow(3, int64(n3)))
}

func pow(a, b int64) int64 {
	ret := int64(1)
	base := a
	for b != 0 {
		if b&1 != 0 {
			ret *= base
			ret %= 1e9 + 7
		}
		base *= base
		base %= 1e9 + 7
	}
	return ret
}
