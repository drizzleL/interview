package main

func statisticalResult(arrayA []int) []int {
	sum := 1
	var zeroCnt int
	for _, num := range arrayA {
		if num == 0 {
			zeroCnt++
			continue
		}
		sum *= num
	}
	ret := make([]int, len(arrayA))
	if zeroCnt >= 2 {
		return ret
	}
	for i, num := range arrayA {
		if zeroCnt > 0 {
			if num == 0 {
				ret[i] = sum
			}
		} else {
			ret[i] = sum / num
		}
	}
	return ret
}
