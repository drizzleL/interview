package main

func sockCollocation(sockets []int) []int {
	var xor int
	for _, s := range sockets {
		xor ^= s
	}
	v := 1
	for v&xor == 0 {
		v <<= 1
	}
	var a, b int
	for _, s := range sockets {
		if s&v == 0 {
			a ^= s
		} else {
			b ^= s
		}
	}
	return []int{a, b}
}
