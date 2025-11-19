package main

func numberOfSubstrings3(s string) int {
	lastZero := make([]int, len(s)+1)
	lastZero[0] = -1
	for i := 1; i <= len(s); i++ {
		lastZero[i] = lastZero[i-1]
		if s[i-1] == '0' {
			lastZero[i] = i - 1
		}
	}
	var ret int
	for i := 0; i < len(s); i++ {
		ret += i - lastZero[i+1]
		currentZeroPos := lastZero[i+1]
		for zeroCnt := 1; zeroCnt*zeroCnt <= i+1 && currentZeroPos != -1; zeroCnt++ {
			lLowerBound := lastZero[currentZeroPos]
			lUpperBoundFromCondition := i + 1 - (zeroCnt*zeroCnt + zeroCnt)
			finalLUpperBound := min(currentZeroPos, lUpperBoundFromCondition)
			ret += max(0, finalLUpperBound-lLowerBound)
			currentZeroPos = lastZero[currentZeroPos]
		}
	}
	return ret
}
