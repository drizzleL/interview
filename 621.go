package main

func leastInterval(tasks []byte, n int) int {
	dict := make([]int, 26)
	var maxLen, maxCnt int
	for _, t := range tasks {
		dict[t-'A']++
		if dict[t-'A'] > maxLen {
			maxLen = dict[t-'A']
			maxCnt = 0
		}
		if dict[t-'A'] == maxLen {
			maxCnt++
		}
	}
	return max(len(tasks), maxCnt+(maxLen-1)*max(n+1, maxCnt))

}
