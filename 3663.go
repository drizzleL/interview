package main

func getLeastFrequentDigit(n int) int {
	var dict [10]int
	for n != 0 {
		dict[n%10] += 1
		n /= 10
	}
	ret := -1
	for i := range dict {
		if dict[i] == 0 {
			continue
		}
		if ret == -1 || dict[i] < dict[ret] || (dict[i] == dict[ret] && i < ret) {
			ret = i
		}
	}
	return ret
}
