package main

func maxDistance2(arrays [][]int) int {
	var ret int
	minVal, maxVal := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		line := arrays[i]
		lineMax, lineMin := line[len(line)-1], line[0]
		ret = max(ret, lineMax-minVal)
		ret = max(ret, maxVal-lineMin)
		minVal = min(minVal, lineMin)
		maxVal = max(maxVal, lineMax)
	}
	return ret
}
