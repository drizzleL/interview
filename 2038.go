package main

func winnerOfGame(colors string) bool {
	var cnt1, cnt2 int
	for i := 0; i < len(colors); i++ {
		j := i
		for j+1 < len(colors) {
			if colors[j+1] != colors[i] {
				break
			}
			j++
		}
		if j-i-1 >= 0 {
			if colors[i] == 'A' {
				cnt1 += j - i - 1
			} else {
				cnt2 += j - i - 1
			}
		}
		i = j
	}
	return cnt1-1 >= cnt2
}
