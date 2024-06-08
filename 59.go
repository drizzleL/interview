package main

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	ret[0][0] = 1
	up, down := 0, n-1
	l, r := 0, n-1
	var i, j int
	v := 1
	for v < n*n {
		for j < r { // go right
			j++
			v++
			ret[i][j] = v
			if j == r {
				up++
				break
			}
		}
		for i < down { // go down
			i++
			v++
			ret[i][j] = v
			if i == down {
				r--
				break
			}
		}
		for j > l { // go left
			j--
			v++
			ret[i][j] = v
			if j == l {
				down--
				break
			}
		}
		for i > up { // go up
			i--
			v++
			ret[i][j] = v
			if i == up {
				l++
				break
			}

		}
	}
	return ret
}
