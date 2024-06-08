package main

func merge(A []int, m int, B []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		var val int
		switch {
		case m != 0 && n != 0:
			a, b := A[m-1], B[n-1]
			if a >= b {
				m--
				val = a
			} else {
				n--
				val = b
			}
		case m == 0:
			n--
			val = B[n]
		case n == 0:
			m--
			val = A[m]
		}
		A[i] = val
	}
}
