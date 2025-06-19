package main

func countPairs7(coordinates [][]int, k int) int {
	dict := map[[2]int]int{}
	var ret int
	for _, cor := range coordinates {
		for a := 0; a <= k; a++ {
			b := k - a
			ret += dict[[2]int{cor[0] ^ a, cor[1] ^ b}]
		}
		dict[[2]int{cor[0], cor[1]}] += 1
	}
	return ret
}
