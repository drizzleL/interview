package main

import "math"

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
	return ret * int(math.Pow(3, float64(n3))) % 1e7
}
