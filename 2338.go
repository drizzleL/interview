package main

func idealArrays(n int, maxValue int) int {
	cnts := make([][]int, maxValue+1) // ways distinct arr + end with val
	for i := range cnts {
		cnts[i] = make([]int, min(14, n+1))
	}
	for i := 1; i <= maxValue; i++ {
		cnts[i][0] += 1
		for j := 2; i*j <= maxValue; j++ {
			for bars := 0; cnts[i][bars] != 0; bars++ {
				cnts[i*j][bars+1] += cnts[i][bars]
			}
		}
	}

	var ret int
	for i := 1; i <= maxValue; i++ {
		for bars := 0; bars < min(14, n); bars++ {
			ret += cnts[i][bars] * combination(n-1, bars)
			ret %= 1e9 + 7
		}
	}
	return ret
}
