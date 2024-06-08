package main

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	tuplesToList := func(tuples [][]int) []int {
		parents := map[int]int{}
		dict := map[int][]int{}
		for _, p := range tuples {
			parents[p[1]] += 1
			dict[p[0]] = append(dict[p[0]], p[1])
		}
		var cc []int
		for i := 1; i <= k; i++ {
			if parents[i] == 0 {
				cc = append(cc, i)
			}
		}
		var ret []int
		for len(cc) != 0 {
			var nextcc []int
			for _, c := range cc {
				ret = append(ret, c)
				for _, child := range dict[c] {
					parents[child]--
					if parents[child] == 0 {
						nextcc = append(nextcc, child)
					}
				}
			}
			cc = nextcc
		}
		return ret
	}
	rows, cols := tuplesToList(rowConditions), tuplesToList(colConditions)
	if len(rows) != k || len(cols) != k {
		return nil
	}
	ret := make([][]int, k)
	for i := range ret {
		ret[i] = make([]int, k)
	}
	colDict := map[int]int{}
	for i, c := range cols {
		colDict[c] = i
	}

	for i, r := range rows {
		ret[i][colDict[r]] = r
	}
	return ret
}
