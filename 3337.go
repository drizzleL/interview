package main

func lengthAfterTransformations2(s string, t int, nums []int) int {
	mat := make([][]int, 26)
	for i := range mat {
		mat[i] = make([]int, 26)
	}
	for i := 0; i < 26; i++ {
		for m := 1; m <= nums[i]; m++ {
			j := (i + m) % 26
			mat[i][j] = 1
		}
	}
	mat2 := matrixPow(mat, t)
	dict := make([]int, 26)
	for _, c := range s {
		dict[c-'a'] += 1
	}
	var ret int
	for i, cnt := range dict {
		for _, c := range mat2[i] {
			ret += cnt * c
			ret %= 1e9 + 7
		}
	}
	return ret
}
func matrixPow(mat [][]int, pow int) [][]int {
	ret := make([][]int, len(mat))
	for i := range ret {
		ret[i] = make([]int, len(mat))
		ret[i][i] = 1
	}
	for pow > 0 {
		if pow&1 != 0 {
			ret = matrixMul(ret, mat)
		}
		pow >>= 1
		mat = matrixMul(mat, mat)
	}
	return ret
}

func matrixMul(mat1, mat2 [][]int) [][]int {
	ret := make([][]int, len(mat1))
	for i := range ret {
		ret[i] = make([]int, len(mat2[0]))
	}
	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret[0]); j++ {
			for m := 0; m < len(mat1[0]); m++ {
				ret[i][j] += mat1[i][m] * mat2[m][j]
				ret[i][j] %= 1e9 + 7
			}
		}
	}
	return ret
}
