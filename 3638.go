package main

func maxBalancedShipments(weight []int) int {
	try := func(i int) int {
		for i += 1; i < len(weight); i++ {
			if weight[i] >= weight[i-1] {
				continue
			}
			break
		}
		return i
	}
	var ret int
	for i := 0; i < len(weight); i++ {
		i = try(i)
		if i == len(weight) {
			break
		}
		ret += 1
	}
	return ret
}
