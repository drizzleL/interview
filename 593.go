package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	same := func(p1, p2 []int) bool {
		return p1[0] == p2[0] && p1[1] == p2[1]
	}
	sum := func(p1, p2 []int) [2]int {
		return [2]int{p1[0] + p2[0], p1[1] + p2[1]}
	}
	gap := func(p1, p2 []int) int {
		a, b := p1[0]-p2[0], p1[1]-p2[1]
		return a*a + b*b
	}
	check := func(p1, p2, p3, p4 []int) bool {
		if same(p1, p2) || same(p3, p4) {
			return false
		}
		if sum(p1, p2) != sum(p3, p4) {
			return false
		}
		return gap(p1, p2) == gap(p3, p4) && gap(p1, p3) == gap(p1, p4)
	}
	return check(p1, p2, p3, p4) || check(p1, p3, p2, p4) || check(p1, p4, p2, p3)
}
