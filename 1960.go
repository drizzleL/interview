package main

func maxProduct(s string) int64 {
	center := make([]int, len(s))
	var c, r int
	for i := 0; i < len(s); i++ {
		mirror := c*2 - i
		if i < r {
			center[i] = min(r-i, center[mirror])
		}
		for a, b := i-center[i]-1, i+center[i]+1; a >= 0 && b < len(s) && s[a] == s[b]; a, b = a-1, b+1 {
			center[i] += 1
		}
		if i+center[i] > r {
			r = i + center[i]
			c = i
		}
	}
	var q [][2]int
	rightMax := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for len(q) != 0 && q[0][1] > i {
			q = q[1:]
		}
		rightMax[i] = 1
		if len(q) != 0 {
			rightMax[i] = 1 + (q[0][0]-i)*2
		}
		if i+1 < len(s) {
			rightMax[i] = max(rightMax[i], rightMax[i+1])
		}
		q = append(q, [2]int{i, i - center[i]})
	}
	var q2 [][2]int
	leftMax := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		for len(q2) != 0 && q2[0][1] < i {
			q2 = q2[1:]
		}
		leftMax[i] = 1
		if len(q2) != 0 {
			leftMax[i] = 1 + (i-q2[0][0])*2
		}
		if i-1 >= 0 {
			leftMax[i] = max(leftMax[i], leftMax[i-1])
		}
		q2 = append(q2, [2]int{i, i + center[i]})
	}
	var ret int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			ret = max(ret, leftMax[i]*rightMax[i+1])
		}
	}
	return int64(ret)
}
