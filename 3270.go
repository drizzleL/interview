package main

func generateKey(num1 int, num2 int, num3 int) int {
	vals := []int{num1, num2, num3}
	var ret int
	for i, base := 0, 1; i < 4; i++ {
		v := vals[0] % 10
		for i, val := range vals {
			v = min(v, val%10)
			vals[i] = val / 10
		}
		ret += base * v
		base *= 10
	}
	return ret
}
