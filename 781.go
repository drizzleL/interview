package main

func numRabbits(answers []int) int {
	dict := map[int]int{}
	var ret int
	for _, ans := range answers {
		dict[ans+1] += 1
	}
	for k, v := range dict {
		ret += (v + k - 1) / k * k
	}
	return ret
}
