package main

func smallestNumber5(pattern string) string {
	revBytes := func(x []byte) {
		for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
			x[i], x[j] = x[j], x[i]
		}
	}
	var ret []byte
	var stack []byte
	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, '1'+byte(i))
		if i == len(pattern) || pattern[i] == 'I' {
			revBytes(stack)
			ret = append(ret, stack...)
			stack = stack[:0]
		}
	}
	return string(ret)
}
