package main

func minimizeXor(num1 int, num2 int) int {
	var bits []int
	if num1 == 0 {
		bits = append(bits, 0)
	}
	for num1 != 0 {
		bits = append(bits, num1&1)
		num1 >>= 1
	}
	var bitSize int
	for num2 != 0 {
		if num2&1 != 0 {
			bitSize += 1
		}
		num2 >>= 1
	}
	ret := make([]int, len(bits))
	for j := len(bits) - 1; j >= 0 && bitSize > 0; j-- {
		if bits[j] == 1 {
			ret[j] = 1
			bitSize -= 1
		}
	}
	for j := 0; j < len(bits) && bitSize > 0; j++ {
		if bits[j] != 1 {
			ret[j] = 1
			bitSize -= 1
		}
	}
	for ; bitSize > 0; bitSize -= 1 {
		ret = append(ret, 1)
	}
	var final int
	for i := 0; i < len(ret); i++ {
		final |= ret[i] * (1 << i)
	}
	return final
}
