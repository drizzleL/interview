package main

func totalStrength(strength []int) int {
	left, right := make([]int, len(strength)), make([]int, len(strength))
	for i := range left {
		left[i] = -1
		right[i] = len(strength)
	}
	var q []int
	sum1, sum2 := make([]int, len(strength)+1), make([]int, len(strength)+1)
	leftAcc, rightAcc := make([]int, len(strength)+1), make([]int, len(strength)+1)
	for i := 0; i < len(strength); i++ {
		for len(q) != 0 && strength[q[len(q)-1]] > strength[i] {
			right[q[len(q)-1]] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
		sum1[i+1] = sum1[i] + strength[i]
		sum1[i+1] %= 1e9 + 7
		leftAcc[i+1] = leftAcc[i] + sum1[i+1]
		leftAcc[i+1] %= 1e9 + 7
	}
	queryLeft := func(a, b int) int {
		return leftAcc[b+1] - leftAcc[a] - sum1[a]*(b-a+1)
	}
	q = q[:0]
	for i := len(strength) - 1; i >= 0; i-- {
		for len(q) != 0 && strength[q[len(q)-1]] >= strength[i] {
			left[q[len(q)-1]] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
		sum2[i] = sum2[i+1] + strength[i]
		sum2[i] %= 1e9 + 7
		rightAcc[i] = rightAcc[i+1] + sum2[i]
		rightAcc[i] %= 1e9 + 7
	}
	queryRight := func(a, b int) int {
		if a > b {
			return 0
		}
		return rightAcc[a] - rightAcc[b+1] - sum2[b+1]*(b-a+1)
	}
	var ret int
	for i := 0; i < len(strength); i++ {
		l, r := left[i], right[i]
		ln, rn := i-l, r-i
		racc := queryLeft(i, r-1)
		lacc := queryRight(l+1, i-1)
		if lacc < 0 {
			lacc += 1e9 + 7
		}
		if racc < 0 {
			racc += 1e9 + 7
		}
		ret += strength[i] * (ln*racc + rn*lacc)
		ret %= 1e9 + 7
	}
	return ret
}
