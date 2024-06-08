package main

func largestInteger(num int) int {
	dict := [10]int{}
	for tmp := num; tmp > 0; tmp /= 10 {
		dict[tmp%10] += 1
	}
	var ret int
	base := 1
	for tmp := num; tmp > 0; tmp /= 10 {
		d := tmp % 10
		var v int
		if d%2 == 0 {
			v = 0
		} else {
			v = 1
		}
		for dict[v] == 0 {
			v += 2
		}
		dict[v] -= 1
		ret += v * base
		base *= 10
	}
	return ret
}
