package main

func sampleStats(count []int) []float64 {
	var sum, cnt, mode int
	minVal, maxVal := -1, -1
	for i, c := range count {
		if c != 0 && minVal == -1 {
			minVal = i
		}
		if c != 0 {
			maxVal = i
		}
		if c > count[mode] {
			mode = i
		}
		sum += i * c
		cnt += c
	}
	var mean float64
	var i int
	for preCnt := 0; preCnt < cnt/2; i++ {
		preCnt += count[i]
		if preCnt > cnt/2 {
			mean = float64(i)
			break
		}
		if preCnt == cnt/2 {
			j := i + 1
			for ; preCnt == cnt/2; j++ {
				preCnt += count[j]
			}
			if cnt%2 == 1 {
				mean = float64(j - 1)
			} else {
				mean = float64(i+j-1) / 2
			}
		}
	}
	ret := make([]float64, 5)
	ret[0] = float64(minVal)
	ret[1] = float64(maxVal)
	ret[2] = float64(sum) / float64(cnt)
	ret[3] = mean
	ret[4] = float64(mode)
	return ret
}
