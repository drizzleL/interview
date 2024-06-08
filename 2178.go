package main

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return nil
	}
	var ret []int64
	var sum int
	var last int
	for last = 2; ; last += 2 {
		sum += last
		if sum >= int(finalSum) {
			break
		}
	}
	for x := 2; x <= last; x += 2 {
		if x == sum-int(finalSum) {
			continue
		}
		ret = append(ret, int64(x))
	}
	return ret
}
