package main

func valueAfterKSeconds(n int, k int) int {
	vals := make([]int, n)
	for i := range vals {
		vals[i] = 1
	}
	for ; k > 0; k-- {
		for i := 1; i < len(vals); i++ {
			vals[i] += vals[i-1]
			vals[i] %= 1e9 + 7
		}
	}
	return vals[len(vals)-1]
}
