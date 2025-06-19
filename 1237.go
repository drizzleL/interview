package main

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var ret [][]int
	for x := 1; customFunction(x, 1) <= z; x++ {
		for y := 1; customFunction(x, y) <= z; y++ {
			if customFunction(x, y) == z {
				ret = append(ret, []int{x, y})
				break
			}
		}
	}
	return ret
}
