package main

func findThePrefixCommonArray(A []int, B []int) []int {
	ret := make([]int, len(A))
	dict := make([]bool, len(A)+1)
	var sum int
	for i := range A {
		a, b := A[i], B[i]
		if dict[a] {
			sum += 1
		}
		dict[a] = true
		if dict[b] {
			sum += 1
		}
		dict[b] = true
		ret[i] = sum
	}
	return ret
}
