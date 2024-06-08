package main

func timeRequiredToBuy(tickets []int, k int) int {
	var ret int
	for i := range tickets {
		if i <= k {
			ret += min(tickets[i], tickets[k])
		} else {
			ret += min(tickets[i], tickets[k]-1)
		}
	}
	return ret
}
