package main

func findSwapValues(array1 []int, array2 []int) []int {
	var sum1, sum2 int
	dict := map[int]bool{}
	for _, v := range array1 {
		sum1 += v
		dict[v] = true
	}
	for _, v := range array2 {
		sum2 += v
	}
	gap := sum1 - sum2
	if gap%2 != 0 {
		return nil
	}
	for _, b := range array2 {
		a := b + gap/2
		if dict[a] {
			return []int{a, b}
		}
	}
	return nil
}
