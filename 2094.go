package main

func findEvenNumbers(digits []int) []int {
	var ret []int
	var dict [10]int
	for _, d := range digits {
		dict[d] += 1
	}
	for i := 100; i < 999; i += 2 {
		a, c := i/100, i%100
		b := i - (a*100-c)/10
		var d [10]int
		d[a] += 1
		d[b] += 1
		d[c] += 1
		var flag bool
		for k := range d {
			if d[k] > dict[k] {
				flag = true
				break
			}
		}
		if !flag {
			ret = append(ret, i)
		}
	}
	return ret
}
