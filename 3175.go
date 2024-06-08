package main

func findWinningPlayer(skills []int, k int) int {
	var idx, wins int
	topSkill := skills[0]
	for i := 1; i < len(skills); i++ {
		sk := skills[i]
		if sk > topSkill {
			wins = 1
			idx = i
			topSkill = sk
		} else {
			wins += 1
		}
		if wins == k {
			return idx
		}
	}
	return idx
}
