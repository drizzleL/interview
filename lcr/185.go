package main

func statisticsProbability(num int) []float64 {
	if num == 0 {
		return nil
	}
	p := float64(1) / float64(6)
	basic := []float64{p, p, p, p, p, p}
	ret := append([]float64{}, basic...)
	for i := 1; i < num; i++ {
		tmp := make([]float64, len(ret)+5)
		for j := 0; j < len(ret); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += ret[j] * basic[k]
			}
		}
		ret = tmp
	}
	return ret
}
