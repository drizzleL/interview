package main

func trainingPlan(actions []int) int {
	dict := [32]int{}
	for _, act := range actions {
		for i := range dict {
			if (1<<i)&act != 0 {
				dict[i]++
			}
		}
	}
	var ret int
	for i := range dict {
		if dict[i]%3 == 0 {
			continue
		}
		ret |= 1 << i
	}
	return ret
}
