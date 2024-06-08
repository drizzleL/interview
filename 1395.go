package main

func numTeams(rating []int) int {
	less, greater := make([]int, len(rating)), make([]int, len(rating))
	for i, num := range rating {
		for j := i + 1; j < len(rating); j++ {
			if rating[j] < num {
				less[i] += 1
			} else {
				greater[i] += 1
			}
		}
	}
	var ret int
	for i := range rating {
		for j := i + 1; j < len(rating); j++ {
			if rating[j] > rating[i] {
				ret += greater[j]
			} else {
				ret += less[j]
			}
		}
	}
	return ret
}
