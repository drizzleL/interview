package main

func countPermutations(complexity []int) int {
	for i := 1; i < len(complexity); i++ {
		if complexity[i] >= complexity[0] {
			return 0
		}
	}
	return fact(len(complexity) - 1)
}
