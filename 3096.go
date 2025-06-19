package main

func minimumLevels(possible []int) int {
	var alice, bob int
	for _, p := range possible {
		bob += p*2 - 1
	}
	for i := 0; i < len(possible)-1; i++ {
		alice += possible[i]*2 - 1
		bob -= possible[i]*2 - 1
		if alice > bob {
			return i + 1
		}
	}
	return -1
}
