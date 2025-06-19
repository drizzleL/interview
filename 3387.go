package main

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	helper := func(pairs [][]string, rates []float64) map[string]float64 {
		fromDict := map[string]map[string]float64{}
		setFrom := func(a, b string, rate float64) {
			if fromDict[a] == nil {
				fromDict[a] = make(map[string]float64)
			}
			fromDict[a][b] = rate
		}
		for i := range pairs {
			a, b := pairs[i][0], pairs[i][1]
			setFrom(a, b, rates[i])
			setFrom(b, a, 1/rates[i])
		}
		day := map[string]float64{
			initialCurrency: 1,
		}
		type elem struct {
			curr string
			num  float64
		}
		eles := []elem{{initialCurrency, 1}}
		for len(eles) > 0 {
			var next []elem
			for _, ele := range eles {
				for k, v := range fromDict[ele.curr] {
					if _, ok := day[k]; ok {
						continue
					}
					day[k] = ele.num * v
					next = append(next, elem{k, day[k]})
				}
			}
			eles = next
		}
		return day
	}
	day1 := helper(pairs1, rates1)
	day2 := helper(pairs2, rates2)
	var ret float64
	for k, v := range day1 {
		if _, ok := day2[k]; !ok {
			continue
		}
		tmp := v * (1 / day2[k])
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
