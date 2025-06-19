package main

func maximumSumOfHeights(maxHeights []int) int64 {
	left := make([]int, len(maxHeights))
	right := make([]int, len(maxHeights))
	var leftH, rightH [][2]int
	var ret int
	for i := 0; i < len(maxHeights); i++ {
		for len(leftH) > 0 && leftH[len(leftH)-1][1] > maxHeights[i] { // pop
			leftH = leftH[:len(leftH)-1]
		}
		leftH = append(leftH, [2]int{i, maxHeights[i]})
		if len(leftH) == 1 {
			left[i] = maxHeights[i] * (i + 1)
			continue
		}
		last := leftH[len(leftH)-2]
		left[i] = left[last[0]] + (i-last[0])*maxHeights[i]
	}
	for i := len(maxHeights) - 1; i >= 0; i-- {
		for len(rightH) > 0 && rightH[len(rightH)-1][1] > maxHeights[i] { // pop
			rightH = rightH[:len(rightH)-1]
		}
		rightH = append(rightH, [2]int{i, maxHeights[i]})
		if len(rightH) == 1 {
			right[i] = maxHeights[i] * (len(maxHeights) - i)
			continue
		}
		last := rightH[len(rightH)-2]
		right[i] = right[last[0]] + (last[0]-i)*maxHeights[i]
	}
	for i := range maxHeights {
		ret = max(ret, left[i]+right[i]-maxHeights[i])
	}
	return int64(ret)
}
