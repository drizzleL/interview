package main

func countCollisions(directions string) int {
	var ret int
	var flag bool
	var cnt int
	for i := 0; i < len(directions); i++ {
		switch directions[i] {
		case 'L':
			if !flag {
				continue
			}
			ret += cnt + 1
			cnt = 0
		case 'R':
			cnt += 1
			flag = true
		case 'S':
			ret += cnt
			cnt = 0
			flag = true
		}
	}
	return ret
}
