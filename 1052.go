package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var sum int
	for i := range customers {
		if grumpy[i] == 1 {
			continue
		}
		sum += customers[i]
	}
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 0 {
			continue
		}
		sum += customers[i]
	}
	ret := sum
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			sum += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			sum -= customers[i-minutes]
		}
		ret = max(ret, sum)
	}
	return ret
}
