package main

func numberOfBoomerangs(points [][]int) int {
	var ret int
	for i, p1 := range points {
		dict := map[int]int{}
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			p2 := points[j]
			a, b := p1[0]-p2[0], p1[1]-p2[1]
			dict[a*a+b*b] += 1
		}
		for _, v := range dict {
			ret += v * (v - 1)
		}
	}
	return ret
}
