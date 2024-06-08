package main

func processQueries(queries []int, m int) []int {
	dict := map[int]int{}
	vals := []int{}
	for i := 1; i <= m; i++ {
		vals = append(vals, i)
		dict[i] = i - 1
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		ret[i] = dict[q]
		for j := dict[q] - 1; j >= 0; j-- {
			val := vals[j]
			dict[val] += 1
			vals[j+1] = val
		}
		dict[q] = 0
		vals[0] = q
	}
	return ret
}
