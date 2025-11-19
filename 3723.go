package main

func maxSumOfSquares(num int, sum int) string {
	if sum > num*9 {
		return ""
	}
	var ret []byte
	for sum > 0 {
		val := min(sum, 9)
		sum -= val
		ret = append(ret, byte(val+'0'))
		num -= 1
	}
	for i := 0; i < num; i++ {
		ret = append(ret, '0')
	}
	return string(ret)
}
