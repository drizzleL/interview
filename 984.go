package main

func strWithout3a3b(a int, b int) string {
	goA := a > b
	var bb []byte
	for ; a != 0 && b != 0; goA = !goA {
		if goA {
			if a-b > 1 {
				bb = append(bb, 'a', 'a')
				a -= 2
			} else {
				bb = append(bb, 'a')
				a -= 1
			}
		} else {
			if b-a > 1 {
				bb = append(bb, 'b', 'b')
				b -= 2
			} else {
				bb = append(bb, 'a')
				b -= 1
			}
		}
	}
	return string(bb)
}
