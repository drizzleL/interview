package main

func countTrapezoids2(points [][]int) int {
	dict := map[[2]int]map[[2]int]int{}
	vertDict := map[int]int{}
	toKey := func(a, b int) [2]int {
		if a == 0 || b == 0 {
			return [2]int{0, 0}
		}
		flag := 1
		if a < 0 {
			a = -a
			flag *= -1
		}
		if b < 0 {
			b = -b
			flag *= -1
		}
		m := gcd(a, b)
		ret := [2]int{a / m, b / m}
		ret[0] *= flag
		return ret
	}
	helper := func(k [2]int, x, y int) [2]int {
		if k[0] == 0 {
			return [2]int{y, 0}
		}
		k[0] *= x
		k[0] -= y * k[1]
		return toKey(k[0], k[1])
	}
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			if p1[0] == p2[0] {
				vertDict[p1[0]] += 1
				continue
			}
			key := toKey(p2[1]-p1[1], p2[0]-p1[0])
			if _, ok := dict[key]; !ok {
				dict[key] = map[[2]int]int{}
			}
			dict[key][helper(key, p1[0], p1[1])] += 1
		}
	}
	midPointCounts := map[[2]int]map[[2]int]int{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			midX := points[i][0] + points[j][0]
			midY := points[i][1] + points[j][1]
			if _, ok := midPointCounts[[2]int{midX, midY}]; !ok {
				midPointCounts[[2]int{midX, midY}] = map[[2]int]int{}
			}
			p1, p2 := points[i], points[j]
			var key [2]int
			if p1[0] == p2[0] {
				key = [2]int{0, -1}
			} else {
				key = toKey(p2[1]-p1[1], p2[0]-p1[0])
			}
			midPointCounts[[2]int{midX, midY}][key] += 1
		}
	}
	var parallelograms int
	for _, count := range midPointCounts {
		var sum int
		for _, v := range count {
			parallelograms += sum * v
			sum += v
		}
	}
	var ret int
	for _, lines := range dict {
		var sum int
		for _, l := range lines {
			ret += l * sum
			sum += l
		}
	}
	var vertSum int
	for _, v := range vertDict {
		ret += v * vertSum
		vertSum += v
	}
	return ret - parallelograms
}
