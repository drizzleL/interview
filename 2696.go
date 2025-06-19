package main

func minLength(s string) int {
	var q []rune
	for _, c := range s {
		q = append(q, c)
		for len(q) >= 2 && ((q[len(q)-2] == 'A' && q[len(q)-1] == 'B') || (q[len(q)-2] == 'C' && q[len(q)-1] == 'D')) {
			q = q[:len(q)-2]
		}
	}
	return len(q)
}
