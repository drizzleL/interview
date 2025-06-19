package main

func maxPoints3(points [][]int) int {
	var ret int
	helper := func(p, p2 []int) [2]float64 {
		k := float64(p[1]-p2[1]) / float64(p[0]-p2[0])
		b := float64(p[1]) - k*float64(p[0])
		return [2]float64{k, b}
	}
	for i := 0; i < len(points); i++ {
		verti := 1
		lineDict := map[[2]float64]int{}
		p := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			if p[0] == p2[0] {
				verti += 1
				ret = max(ret, verti)
				continue
			}
			key := helper(p, p2)
			if lineDict[key] == 0 { // init
				lineDict[key] = 2
			}
			lineDict[key] += 1
			ret = max(ret, lineDict[key])
		}
	}
	return ret
}
