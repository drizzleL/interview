package main

func countEven(num int) int {
	baseNum := num / 10 * 10
	ret := max(0, baseNum/2-1)
	var flag bool
	for m := baseNum; m != 0; m /= 10 {
		if m%10%2 == 1 {
			flag = !flag
		}
	}
	if flag {
		baseNum += 1
	}
	baseNum = max(2, baseNum)
	if num >= baseNum {
		ret += (num-baseNum)/2 + 1
	}
	return ret
}
