package main

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	var presum, ret int
	for i := 0; i <= len(startTime); i++ {
		switch i {
		case 0:
			presum += startTime[0]
		case len(startTime):
			presum += eventTime - endTime[len(endTime)-1]
		default:
			presum += startTime[i] - endTime[i-1]
		}
		if i > k { // gotta remove before
			if i-k == 1 {
				presum -= startTime[0]
			} else {
				presum -= startTime[i-k-1] - endTime[i-k-2]
			}
		}
		ret = max(ret, presum)
	}
	return ret
}
