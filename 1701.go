package main

func averageWaitingTime(customers [][]int) float64 {
	if len(customers) == 0 {
		return 0
	}
	var sum int
	now := customers[0][0]
	for _, c := range customers {
		if c[0] < now {
			sum += now - c[0]
		} else {
			now = c[0]
		}
		sum += c[1]
		now += c[1]
	}
	return float64(sum) / float64(len(customers))
}
