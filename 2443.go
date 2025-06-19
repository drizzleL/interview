package main

func sumOfNumberAndReverse(num int) bool {
	reverse := func(x int) int {
		var ret int
		for x != 0 {
			ret = ret*10 + x%10
			x /= 10
		}
		return ret
	}
	for i := 0; i <= num; i++ {
		if num == reverse(i)+i {
			return true
		}
	}
	return false
}
