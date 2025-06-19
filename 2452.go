package main

func twoEditWords(queries []string, dictionary []string) []string {
	var ret []string
	for _, q := range queries {
		for _, w := range dictionary {
			var cnt int
			for i := range q {
				if q[i] == w[i] {
					continue
				}
				cnt += 1
			}
			if cnt <= 2 {
				ret = append(ret, q)
				break
			}
		}
	}
	return ret
}
