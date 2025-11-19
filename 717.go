package main

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			if i == len(bits)-1 {
				return true
			}
			continue
		}
		i += 1
	}
	return false
}
