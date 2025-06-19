package main

func smallestNumber4(num int64) int64 {
	var flag bool
	if num < 0 {
		flag = true
		num = -num
	}
	var dict [10]int
	for num != 0 {
		dict[num%10] += 1
		num /= 10
	}
	var ret int
	if flag {
		for i := 9; i >= 0; i-- {
			for j := 0; j < dict[i]; j++ {
				ret *= 10
				ret += i
			}
		}
		return int64(-ret)
	} else {
		for i := 1; i <= 9; i++ {
			if dict[i] != 0 {
				dict[i] -= 1
				ret = i
				break
			}
		}
		for i := 0; i <= 9; i++ {
			for j := 0; j < dict[i]; j++ {
				ret *= 10
				ret += i
			}
		}
	}
	return int64(ret)
}
