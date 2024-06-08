package main

func minimumSum(num int) int {
	dict := [10]int{}
	for num != 0 {
		dict[num%10] += 1
		num /= 10
	}
	var ret int
	for i, n := 0, 0; i < 10; i++ {
		for dict[i] != 0 && n < 2 {
			ret += i * 10
			dict[i] -= 1
			n += 1
		}
	}
	for i := 9; i >= 0; i-- {
		for dict[i] != 0 {
			ret += i
			dict[i] -= 1
		}
	}
	return ret
}
