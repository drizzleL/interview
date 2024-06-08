package main

func bestLine(points [][]int) []int {
	if len(points) <= 1 {
		return nil
	}
	var ret []int
	check := func(a, b, c []int) bool {
		numerator1 := a[1] - b[1]
		denominator1 := a[0] - b[0]
		numerator2 := a[1] - c[1]
		denominator2 := a[0] - c[0]
		return numerator1*denominator2 == numerator2*denominator1
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			tmp := []int{i, j}
			for k := j + 1; k < len(points); k++ {
				c := points[k]
				if check(a, b, c) {
					tmp = append(tmp, k)
				}
			}
			if len(tmp) > len(ret) {
				ret = tmp
			}
		}
	}
	return ret[:2]
}
