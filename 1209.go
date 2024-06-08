package main

func removeDuplicates(s string, k int) string {
	var stack [][2]int
	for _, c := range s {
		if len(stack) != 0 && stack[len(stack)-1][0] == int(c) {
			stack[len(stack)-1][1] += 1
		} else {
			stack = append(stack, [2]int{int(c), 1})
		}
		for len(stack) != 0 && stack[len(stack)-1][1] >= k {
			stack[len(stack)-1][1] -= k
			if stack[len(stack)-1][1] == 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	var ret []byte
	for _, v := range stack {
		for i := 0; i < v[1]; i++ {
			ret = append(ret, byte(v[0]))
		}
	}
	return string(ret)
}
